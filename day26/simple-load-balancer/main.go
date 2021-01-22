
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
  }

type ServerPool struct {
	backends []*Backend
	current  uint64
  }
  
u, _ := url.Parse("http://localhost:8080")
rp := httputil.NewSingleHostReverseProxy(u)
  
// initialize your server and add this as handler
http.HandlerFunc(rp.ServeHTTP)

func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
  }

// GetNextPeer returns next active peer to take a connection
func (s *ServerPool) GetNextPeer() *Backend {
	// loop entire backends to find out an Alive backend
	next := s.NextIndex()
	l := len(s.backends) + next // start from next and move a full cycle
	for i := next; i < l; i++ {
	  idx := i % len(s.backends) // take an index by modding with length
	  // if we have an alive backend, use it and store if its not the original one
	  if s.backends[idx].IsAlive() {
		if i != next {
		  atomic.StoreUint64(&s.current, uint64(idx)) // mark the current one
		}
		return s.backends[idx]
	  }
	}
	return nil
  }

  // SetAlive for this backend
func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
  }

  // IsAlive returns true when backend is alive
func (b *Backend) IsAlive() (alive bool) {
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return
  }
  
  // lb load balances the incoming request
func lb(w http.ResponseWriter, r *http.Request) {
	peer := serverPool.GetNextPeer()
	if peer != nil {
	  peer.ReverseProxy.ServeHTTP(w, r)
	  return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
  }

  server := http.Server{
	Addr:    fmt.Sprintf(":%d", port),
	Handler: http.HandlerFunc(lb),
  }
  
proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
  log.Printf("[%s] %s\n", serverUrl.Host, e.Error())
  retries := GetRetryFromContext(request)
  if retries < 3 {
    select {
      case <-time.After(10 * time.Millisecond):
        ctx := context.WithValue(request.Context(), Retry, retries+1)
        proxy.ServeHTTP(writer, request.WithContext(ctx))
      }
      return
    }

  // after 3 retries, mark this backend as down
  serverPool.MarkBackendStatus(serverUrl, false)

  // if the same request routing for few attempts with different backends, increase the count
  attempts := GetAttemptsFromContext(request)
  log.Printf("%s(%s) Attempting retry %d\n", request.RemoteAddr, request.URL.Path, attempts)
  ctx := context.WithValue(request.Context(), Attempts, attempts+1)
  lb(writer, request.WithContext(ctx))
}