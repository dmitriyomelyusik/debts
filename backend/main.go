package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dmitriyomelyusik/debts/backend/database"
	"github.com/dmitriyomelyusik/debts/backend/rest"
	"github.com/dmitriyomelyusik/debts/backend/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := mysql.NewDB()
	service := service.NewService(db)
	ctrl := rest.NewController(&service)
	router := rest.NewRouter(&ctrl)
	server := http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      router,
	}
	log.Println("started")
	log.Fatal(server.ListenAndServe())
}
