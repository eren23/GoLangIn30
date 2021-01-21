
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
  