package main

import (
	"log"

	"http-server-projeto-korp/internal/server"
)

func main() {
	srv := server.New()

	log.Println("Servidor iniciado na porta 8080")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}