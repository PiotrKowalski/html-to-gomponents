package main

import (
	"html-to-gomponents/internal/app"
	"html-to-gomponents/internal/router"
	"log"
)

func main() {
	application := app.New(app.Config{})

	rout, err := router.New(application)
	if err != nil {
		log.Fatal(err)
	}
	rout.Logger.Fatal(rout.Start("127.0.0.1:3050"))
}
