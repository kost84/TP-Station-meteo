package main

import (
	"log"
	"net/http"
)

func main(){

	//helper.LoadTemplate()
	//router := routes.MainRouter 

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080"))
}