package main

import (
	"fmt"

	"github.com/cyn166/urubu-do-pix/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.SetupDB()
	config.SetupRoutes(router)

	router.Run(":8080")

	fmt.Println("Hello, world!")
}
