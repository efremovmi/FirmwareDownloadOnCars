package main

import (
	"Clever_City/internal/service/worker"
	"Clever_City/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/signal"
)

func main() {
	uc := usecase.NewHandler()
	//router := api.NewRouter(uc)
	//
	//go func() {
	//	router.Start()
	//}()

	go func() {
		if err := worker.ProcessingRaspberryFiles(uc); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
