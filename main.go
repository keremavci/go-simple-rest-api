package main

import (
	"net/http"
	"log"
	"github.com/keremavci/go-simple-rest-api/api"
)

func main(){
	log.Print("Starting...")
	log.Fatal(http.ListenAndServe(":9000",api.Handlers()))
}
