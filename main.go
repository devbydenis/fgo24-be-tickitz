package main

import (
	"backend-cinemax/routers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           fgo24-be-tickitz
// @version         1.0
// @description     This is a backend service for tickitz web app
// @Basepath /
func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	routers.CombineRouters(r)
	
	godotenv.Load()
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}