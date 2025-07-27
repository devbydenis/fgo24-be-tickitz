package routers

import (
	c "backend-cinemax/controllers"
	m "backend-cinemax/middlewares"

	"github.com/gin-gonic/gin"
)

func transactionRouters(r *gin.RouterGroup) {
	r.POST("/create", m.AuthMiddleware(), c.CreateTransactionHandler)
	r.GET("/history", m.AuthMiddleware(), c.GetHistoryHandler)
}