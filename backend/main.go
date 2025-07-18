package main

import (
	"log"
	"cadastro-clientes-api/database"
	"cadastro-clientes-api/routes"
)

func main() {
	database.Connect()

	r := routes.SetupRoutes()

	log.Println("Servidor iniciando na porta 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
