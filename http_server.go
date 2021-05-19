package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Url struct {
	Path   string
	Method string
	View   View
}

type HttpServer struct {
	Urls    []Url
	Port    string
	Static  string
	Router  *mux.Router
	Factory HandlerFactory
}

func CreateHttpServer() HttpServer {
	server := HttpServer{
		Urls:    urls,
		Port:    ":" + os.Getenv("PORT"),
		Static:  "./html/static",
		Factory: HandlerFactory{},
	}
	return server
}

func (server HttpServer) Start() {
	server.Router = mux.NewRouter()
	server.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(server.Static))))
	server.registerUrls()
	log.Fatal(http.ListenAndServe(server.Port, nil))
}

func (server *HttpServer) registerUrl(url Url) {
	server.Router.HandleFunc(url.Path, server.Factory.Build(url.View)).Methods(url.Method)
}

func (server *HttpServer) registerUrls() {
	for _, url := range server.Urls {
		server.registerUrl(url)
	}
	http.Handle("/", server.Router)
}
