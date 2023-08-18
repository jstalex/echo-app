package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"metrics/internal/app"
	"metrics/internal/storage"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	name := os.Getenv("PG_NAME")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASSWORD")
	host := os.Getenv("PG_HOST")

	url := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", user, pass, host, name)
	// url := "postgres://admin:admin@localhost:5432/data"
	s, err := storage.Open(ctx, url)
	defer func() {
		err := s.Close(ctx)
		if err != nil {
			log.Printf(err.Error())
		}
	}()

	a, err := app.New(s)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = a.Run(":80")
	if err != nil {
		log.Fatalf(err.Error())
	}
}
