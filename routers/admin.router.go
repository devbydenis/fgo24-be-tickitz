package routers

import (
	c "backend-cinemax/controllers"

	"github.com/gin-gonic/gin"
)

func adminRouters(r *gin.RouterGroup) {
	r.POST("", c.CreateMovieAdminHandler)
	r.GET("/list", c.ListMovieAdminHandler)
	r.PATCH("/update", c.UpdateMovieAdminHandler)
	r.DELETE("/delete/:id", c.DeleteMovieAdminHandler)
}