package main

import (
	"log"
	"net/http"
	"os"

	r "github.com/skantuz/curso/routes"
)

func main() {

	router := r.NewRouter()
	port := os.Getenv("sys_port") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8001" //localhost
	}

	log.Println("Listen" + port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Print(err)
	}
}
