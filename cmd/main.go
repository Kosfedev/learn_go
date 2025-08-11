package main

import (
	"context"
	"log"

	"github.com/Kosfedev/learn_go/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err = a.RunGRPCServer()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = a.RunGRPCGWServer()
	if err != nil {
		log.Fatal(err)
	}
}
