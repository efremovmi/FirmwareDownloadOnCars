package main

import (
	"Clever_City/internal/api"
	"Clever_City/internal/repository"
	"Clever_City/internal/usecase"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "db.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	uc := usecase.NewHandler(repo)
	router := api.NewRouter(uc)

	go func() {
		router.Start()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
