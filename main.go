package main

import (
	"html-to-gomponents/internal/app"
	"html-to-gomponents/internal/router"
	"log"
	"os"
)

func main() {
	var port string
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3060"
	}
	application := app.New(app.Config{})

	rout, err := router.New(application)
	if err != nil {
		log.Fatal(err)
	}
	rout.Logger.Fatal(rout.Start(":" + port))
}
