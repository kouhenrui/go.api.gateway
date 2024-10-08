package main

import (
	"fmt"
	"go.api.gateway/src/config"
	"go.api.gateway/src/router"
	"log"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	fmt.Println(config.Port, "/////////////////")
	//config.Port
	if err := r.Run(config.Port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

}
