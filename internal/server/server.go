package server

import (
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Server interface {
	Serve(addr string)
	Get(path string, handlerFunc HandlerFunc)
	Post(path string, handlerFunc HandlerFunc)
}

type serverImpl struct {
	server *http.Server
}

func NewServer() Server {
	return &serverImpl{
		server: &http.Server{},
	}
}

func (s serverImpl) Get(path string, handlerFunc HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
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
		handlerFunc(w, r)
	})
}

func (s serverImpl) Serve(addr string) {
	s.server.Addr = addr
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
