package routers

import (
	c "backend-cinemax/controllers"

	"github.com/gin-gonic/gin"
)

func authRouters(r *gin.RouterGroup)  {
	r.POST("/register", c.RegisterHandler)
	r.POST("/login", c.LoginHandler)
}