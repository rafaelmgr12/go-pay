package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rafaelmgr12/go-pay/configs"
	"github.com/rafaelmgr12/go-pay/internal/infra/repository"
	"github.com/rafaelmgr12/go-pay/internal/infra/web/handlers"
	webserver "github.com/rafaelmgr12/go-pay/internal/infra/web/server"
	"github.com/rafaelmgr12/go-pay/internal/usecase/transactions"
)

func main() {
	log.Println("Starting application...")

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	conn, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repo := repository.NewChatRepositoryMySQL(conn)
	usecase := transactions.NewTransactionsUseCase(repo)

	webserver := webserver.NewWebServer(":" + configs.WebServerPort)
	webserverBalanceHandler := handlers.NewWebBalanceHandler(*usecase)

	webserver.AddHandler("/api/v1/balance/{_id}", webserverBalanceHandler.GetBalanceHandle)
	webserver.AddHandler("/api/v1/transfer", webserverBalanceHandler.TransferHandle)

	log.Println("Server running on port " + configs.WebServerPort)

	webserver.Start()

}
