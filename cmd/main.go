package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rafaelmgr12/go-pay/configs"
	"github.com/rafaelmgr12/go-pay/internal/infra/web"
)

func main() {
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

	log.Println("Starting application...")
	log.Println("Listening on port " + configs.WebServerPort)

	web.HandleRequests(configs.WebServerPort, conn)
}
