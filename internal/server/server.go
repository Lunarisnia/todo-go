package server

import (
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)
type Middleware HandlerFunc

type Server interface {
	Serve(addr string)
	Get(path string, handlerFunc HandlerFunc)
	Post(path string, handlerFunc HandlerFunc)
	AddMiddleware(middleware Middleware)
}

type serverImpl struct {
	server      *http.Server
	middlewares []Middleware
}

func NewServer() Server {
	return &serverImpl{
		server:      &http.Server{},
		middlewares: make([]Middleware, 0),
	}
}

func (s serverImpl) Get(path string, handlerFunc HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		for _, m := range s.middlewares {
			m(w, r)
		}
		handlerFunc(w, r)
	})
}

func (s serverImpl) Post(path string, handlerFunc HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}
		for _, m := range s.middlewares {
			m(w, r)
		}
		handlerFunc(w, r)
	})
}

func (s *serverImpl) AddMiddleware(middleware Middleware) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s serverImpl) Serve(addr string) {
	s.server.Addr = addr
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
