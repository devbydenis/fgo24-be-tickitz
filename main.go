package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	
	godotenv.Load()
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}