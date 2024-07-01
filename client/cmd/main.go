package main

import (
	"log"
	"net/http"
	"task/client/api"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	mux := api.Routes()

	log.Println("Client is Listening on port", ":7777")
	if err := http.ListenAndServe("localhost:7777", mux); err != nil {
		log.Fatal("Unable to listen and serve :", err)
	}

}
