package main

import (
	"book/app"
	"book/config"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = config.InitPostgres()
	if err != nil {
		panic(err)
	}

}

func main() {
	app.StartApplication()
}
