package main

import (
	"Clever_City/internal/api"
	"Clever_City/internal/repository"
	"Clever_City/internal/usecase"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/ssh"
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

	config := &ssh.ClientConfig{
		User:            "rts-ubuntu-01",
		Auth:            []ssh.AuthMethod{ssh.Password("q1w2e3r4t5")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "192.168.0.104:22", config)
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to dial: %s", err))
	}

	repo := repository.NewRepository(db)
	uc := usecase.NewHandler(repo, client)
	defer uc.Close()
	router := api.NewRouter(uc)

	go func() {
		router.Start()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
