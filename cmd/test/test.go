package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nupamore/lambda/services"
)

func main() {
	godotenv.Load("cmd/test/.env")

	target, err := services.GetRandomImage("test")

	log.Println(*target, err)
}
