package routers

import (
	c "backend-cinemax/controllers"

	"github.com/gin-gonic/gin"
)

func movieRouters(r *gin.RouterGroup) {
	r.GET("/now-showing", c.GetNowShowingMoviesHandler)
	r.GET("/up-coming", c.GetUpComingMoviesHandler)
}