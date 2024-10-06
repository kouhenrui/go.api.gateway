package main

import (
	"go.api.gateway/src/router"
	"log"
)

func main() {
	//config.InitConfig()
	r := router.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

}
