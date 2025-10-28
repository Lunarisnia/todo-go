package server

import (
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

type Server interface {
	Serve(addr string)
	Get(path string, handler Handler)
	Post(path string, handler Handler)
}

type serverImpl struct {
	server *http.Server
}

func NewServer() Server {
	return &serverImpl{
		server: &http.Server{},
	}
}

func (s serverImpl) Get(path string, handler Handler) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	})
}

func (s serverImpl) Post(path string, handler Handler) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	})
}

func (s serverImpl) Serve(addr string) {
	s.server.Addr = addr
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
