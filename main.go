package main

import (
	"github.com/TualatinX/durian-go/initialize"
	"log"
	"os"
)

func main() {

	initialize.InitMySQL()

	r := initialize.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
        log.Fatal("$PORT must be set")
    }
	r.Run(":"+port)
}
