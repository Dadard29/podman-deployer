package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var globalService Service

type Service struct {
	server http.Server
	config config
}

func sendResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	if _, err := w.Write([]byte(message)); err != nil {
		log.Println(err)
	}
}

func NewService() Service {
	config := retrieveConfig()

	addr := fmt.Sprintf("%s:%d", config.host, config.port)
	router := mux.NewRouter()
	router.HandleFunc("/auth", authRoute).Methods(http.MethodPost)
	router.HandleFunc("/deploy", deployRoute).Methods(http.MethodPost)

	server := http.Server{
		Addr:              addr,
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	globalService = Service{
		server: server,
		config: config,
	}

	return globalService
}

func (s Service) Run() {
	log.Printf("listening on %s\n", s.server.Addr)
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
