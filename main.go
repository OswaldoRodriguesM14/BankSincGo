package main

import (
	"context"
	"log"
	"net/http"

	"github.com/OswaldoRodriguesM14/BankSincGo/internal/infra/database"
	"github.com/OswaldoRodriguesM14/BankSincGo/internal/infra/webserver/routes"
)

func main() {

	ConnectionDB, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Não foi possível conectar ao banco de dados: ", err)
	}

	defer func() {
		if err = ConnectionDB.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r := routes.NewRouter(ConnectionDB)

	// Inicia o servidor na porta 5050
	if err := http.ListenAndServe(":5001", r); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
