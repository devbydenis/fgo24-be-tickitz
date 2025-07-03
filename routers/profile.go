package routers

import (
	c "backend-cinemax/controllers"

	"github.com/gin-gonic/gin"
)

func profileRouters(r *gin.RouterGroup) {
	r.GET("/:email", c.GetProfileHandler) 
	// r.PUT("/profile", c.UpdateProfileHandler)
	// r.DELETE("/profile", c.DeleteProfileHandler)
	
}