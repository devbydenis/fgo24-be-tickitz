package routers

import (
	c "backend-cinemax/controllers"

	"github.com/gin-gonic/gin"
)

func movieRouters(r *gin.RouterGroup) {
	r.GET("/now-showing", c.GetNowShowingMoviesHandler) // Endpoint to get now showing movies
	r.GET("/now-playing", c.GetNowShowingMoviesHandler) // Endpoint to get now playing movies
}