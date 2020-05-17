package main

import (
	"log"

	"github.com/frank1995alfredo/twittor-go/bd"
	"github.com/frank1995alfredo/twittor-go/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
