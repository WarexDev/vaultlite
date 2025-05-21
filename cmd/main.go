// @title           VaultLite API
// @version         1.0
// @description     VaultLite is a lightweight secrets management API for secure storage and retrieval of secrets in cloud-native environments.
// @termsOfService  https://github.com/WarexDev/vaultlite/blob/master/LICENSE

// @contact.name   Benoit (WarexDev)
// @contact.url    https://github.com/WarexDev
// @contact.email  benoit.fraysse@hotmail.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

// @schemes http
package main

import (
	_ "github.com/WarexDev/vaultlite/docs"
	"github.com/WarexDev/vaultlite/internal/api"
)

func main() {
	router := api.SetupRouter()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
