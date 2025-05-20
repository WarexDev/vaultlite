package main

import (
	"github.com/WarexDev/vaultlite/internal/api"
)

func main() {
	router := api.SetupRouter()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
