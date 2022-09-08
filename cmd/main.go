package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/maskot/pkg/delivery/jsonrpc"
	"github.com/maskot/pkg/repositories"
	"github.com/maskot/pkg/repositories/postgres"
	"github.com/maskot/pkg/use_cases"
)

func main() {
	db, err := postgres.NewPostgresDB(&postgres.Config{
		DBName:   os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	useCases := use_cases.NewUseCases(&use_cases.Dependencies{Repos: repositories.NewRepository(db)})

	r, err := jsonrpc.NewJSONRpc(&jsonrpc.Dependencies{
		UseCases: useCases,
	})
	if err != nil {
		log.Fatalf("Fail when register api")
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Println(errors.New("Fail to serve jsonrpc"))
			signalChan <- os.Interrupt
		}
	}()

	log.Println("Server started and listening")
	<-signalChan
	gracefulStop(db)
}

func gracefulStop(db *sqlx.DB) {
	log.Println("Closing server...")
	defer log.Println("Server closed!")

	log.Println("Closing PostgreSQL connections")
	if err := db.Close(); err != nil {
		log.Println("fail when closing database")
	}
}
