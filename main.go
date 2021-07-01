package main

import (
	"github.com/TualatinX/durian-go/initialize"
)

func main() {

	initialize.InitMySQL()

	r := initialize.SetupRouter()

	r.Run(":8080")
}
