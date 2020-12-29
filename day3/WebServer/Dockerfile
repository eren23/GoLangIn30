FROM golang:1.15
WORKDIR /src
## COPY go.sum go.mod ./
## RUN go mod download incase of having having additional dependencies we can use that two lines to use cache

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app .
ENTRYPOINT ["/bin/app"]

# FROM alpine
# COPY --from=0 /bin/app /bin/app
# ENTRYPOINT ["/bin/app"]

## Multistage build is to be fixed later, multistage static pages are not viewable.