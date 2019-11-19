package main

import (
	"database/sql"
	"fmt"
	"github.com/encasol/tipsterchat/delivery"
	"github.com/encasol/tipsterchat/repository"
	"github.com/encasol/tipsterchat/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbSource := "tipster:tipster@tcp(mysql)/tipster?charset=utf8"

	connectionSql, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	repo := repository.MySqlTipRepository{Connection: connectionSql}
	svc := service.TipService{TipRepo: repo}
	httpHandler := delivery.HttpTipHandler{TipService: svc}
	fmt.Println("Configuration Done")
	httpHandler.ListenAndServe("localhost", 8080)

}
