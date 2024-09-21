package main

import (
	"back/config"
	"back/routers"
)

func main() {
	config.ConnectDB()
	config.InitRedis()

	router := routers.SetUpRouter()

	router.Run(":8080")

}
