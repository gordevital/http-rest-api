package main

import (
	"log"

	"github.com/alexcesaro/log/stdlog"
	"github.com/gordevital/http-rest-api/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	logger := stdlog.GetFromFlags()
	logger.Info("Server Starting...")
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
