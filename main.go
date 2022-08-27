package main

import (
	"log"

	"github.com/danielsanchez0/goland/handlers"
)

func main() {
	if database.chequeoBase() == 0 {
		log.Fatal("sin conexion a la db")
		return
	}

	handlers.Manejadores()
}
