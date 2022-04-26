package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"meli.test/database"
	"meli.test/server"
)

var successClose bool = false

func main() {
	ctx := context.Background()

	database.Initialize()

	serverDoneChannel := make(chan os.Signal, 1)
	signal.Notify(serverDoneChannel, os.Interrupt, syscall.SIGTERM)

	srv := server.Create(":8080")

	go func() {
		err := srv.ListenAndServe()

		if err != nil && !successClose {
			panic(err)
		}
	}()

	log.Println("server started")

	<-serverDoneChannel

	successClose = true
	srv.Shutdown(ctx)
	log.Println("server stopped")
}
