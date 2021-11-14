package main

import (
	"entertaime-server/config"
	"entertaime-server/routes"
)

func main() {
	config.Connect()
	routes.RunRouter()
}
