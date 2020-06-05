package main

import (
	"github.com/joho/godotenv"

	"github.com/tanimutomo/clean-architecture-api-go/infrastructure"
)

func main() {
	godotenv.Load()
	infrastructure.Router.Run()
}
