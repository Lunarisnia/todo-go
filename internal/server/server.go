package server

import (
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

type Server interface {
	Serve(addr string)
	Register(path string, handler Handler)
}

type serverImpl struct {
	server *http.Server
}

func NewServer() Server {
	return &serverImpl{
		server: &http.Server{},
	}
}

func (s serverImpl) Register(path string, handler Handler) {
	http.HandleFunc(path, handler)
}

func (s serverImpl) Serve(addr string) {
	s.server.Addr = addr
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
