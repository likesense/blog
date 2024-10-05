package main

import (
	"blog/internal/app"
	"log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalf("Error initializing app: %v", err)
	}
	err = a.Run()
	if err != nil {
		log.Fatalf("Error running app: %v", err)
	}
}
