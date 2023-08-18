package main

import (
	"log"

	"metrics/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = a.Run(":80")
	if err != nil {
		log.Fatalf(err.Error())
	}
}
