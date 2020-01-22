package main

import (
	"log"
	"github.com/gordevital/http-rest-api"
)

func main(){
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
