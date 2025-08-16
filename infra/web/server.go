package web

import (
	"net/http"
)

type HttpOption func(s *http.Server)

func NewServer(addr string, handler http.Handler, opts ...HttpOption) *http.Server {
	s := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	for _, f := range opts {
		f(s)
	}

	return s
}
