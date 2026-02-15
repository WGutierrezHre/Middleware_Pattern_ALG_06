package main

import (
	"ejercicio5/server"
	"log"
)

func main() {
	finalHandler := func(request string) {
		log.Println("  Ejecutando: ", request)
	}

	handler := server.Chan(
		finalHandler,
		server.Logging,
		server.Auth,
	)

	log.Println(" --- Prueba con Admin --- ")
	handler("admin")
	log.Println(" --- Prueba sin Admin --- ")
	handler("invitado")

}
