package routers

import (
	c "backend-cinemax/controllers"

	"github.com/gin-gonic/gin"
)

func adminRouters(r *gin.RouterGroup) {
	r.POST("", c.CreateAdminHandler)
	r.GET("/list", c.ListAdminHandler)
	// r.PATCH("/update/:id", c.UpdateAdminHandler)
	// r.DELETE("/delete/:id", c.DeleteAdminHandler)
}