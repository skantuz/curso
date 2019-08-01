package main

import(
	"http"
	"os"
	"github.com/skantuz/backend/routes"
)

func main(){

	route := routes.NewRouter()
	port := os.Getenv("sys_port")
	if ( port == ""){
		port = "8000"
	}
	log.Println("Escuchando en "+port)

	err := http.ListenAndServe(":"+port, route)
	if err!=nil{
		log.Fatal(err)
	}
}