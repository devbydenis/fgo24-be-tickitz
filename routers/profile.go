package routers

import (
	c "backend-cinemax/controllers"
	m "backend-cinemax/middlewares"

	"github.com/gin-gonic/gin"
)

func profileRouters(r *gin.RouterGroup) {
	r.GET("", m.AuthMiddleware(), c.GetProfileHandler) 
	// r.PATCH("/profile", m.AuthMiddleware(), c.UpdateProfileHandler)
	// r.DELETE("/profile", c.DeleteProfileHandler)
	
}