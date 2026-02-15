package server

import (
	"log"
)

var Tracker []string // slice para el test

func ResetTracker() {
	Tracker = []string{}
}

type Handler func(string)

func Logging(next Handler) Handler {
	return func(request string) {
		Tracker = append(Tracker, "Logging")
		log.Println("  Request Recibida: ", request)
		next(request)
		log.Println("  Request Procesada: ", request)
	}
}

func Auth(next Handler) Handler {
	return func(request string) {
		Tracker = append(Tracker, "Auth")
		if request != "admin" {
			log.Println("  Acceso Denegado")
			return
		}
		log.Println("  Usuario Autenticado")
		next(request)
	}
}

func Chan(handler Handler, middlewares ...func(Handler) Handler) Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
