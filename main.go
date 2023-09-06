package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/routers"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	server := routers.InitRouter()
	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	server.Run(port)
}
