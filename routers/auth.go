package routers

import (
	c "backend-cinemax/controllers"
	"backend-cinemax/middlewares"

	"github.com/gin-gonic/gin"
)

func authRouters(r *gin.RouterGroup)  {
	r.POST("/register", c.RegisterHandler)
	r.POST("/login", c.LoginHandler)
	r.POST("/forgot-password", middlewares.AuthMiddleware(), c.ForgotPasswordHandler)
	r.POST("/change-password", middlewares.AuthMiddleware(), c.ChangePasswordHandler)
	r.POST("/verify-otp", middlewares.AuthMiddleware(), c.VerifyOTPHandler)
}