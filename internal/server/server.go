package server

import (
	"net/http"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	server := &Server{
		port: 8081,
	}

	router := server.RegisterRoutes()

	return &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
}
