package routers

import (
	c "backend-cinemax/controllers"

	"github.com/gin-gonic/gin"
)

func authRouters(r *gin.RouterGroup)  {
	r.POST("/register", c.RegisterHandler)
	r.POST("/login", c.LoginHandler)
	r.POST("/forgot-password", c.ForgotPasswordHandler)
	r.POST("/change-password", c.ChangePasswordHandler)
	r.POST("/verify-otp", c.VerifyOTPHandler)
}