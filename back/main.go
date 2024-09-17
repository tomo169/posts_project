package main

import (
	"back/config"
	"back/routers"
)

func main() {
	config.ConnectDB()

	router := routers.SetUpRouter()

	router.Run(":8080")

}
