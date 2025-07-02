package controllers

import (
	"backend-cinemax/config"
	m "backend-cinemax/models"
	u "backend-cinemax/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @summary Handle register user
// @Description Create a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body m.RegisterRequest true "request create user"
// @Success 201 {object} m.RegisterRequest
// @Failure 400 {object} u.Response{Success bool, Message string, Errors any}
// @Router /auth/register [post]
func RegisterHandler(ctx *gin.Context) {
	var req m.RegisterRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})
	}

	if req.Email == "" {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Email is required",
		})
		return
	}

	if req.Password == "" {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Password is required",
		})
		return
	}

	if len(req.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Password must be at least 6 characters",
		})
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Password and Confirm Password did not match",
		})
		return
	}

	if m.IsEmailExist(req.Email) {
		ctx.JSON(http.StatusConflict, u.Response{
			Success: false,
			Message: "Email already exist",
		})
		return
	}
	
	userUUID := u.GenerateUUID()
	parseUserUUID, err := uuid.Parse(userUUID)
	err = m.InsertUserToDB(req.Email, req.Password, parseUserUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, u.Response{
			Success: false,
			Message: "Internal Server Error: Failed to insert user to database",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, u.Response{
		Success: true,
		Message: "Register Success",
	})
}

// @summary Handle login user
// @Description Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body m.LoginRequest true "request login user"
// @Success 200 {object} u.Response{Success bool, Message string, Errors any}
// @Failure 400 {object} u.Response{Success bool, Message string, Errors any}
// @Failure 401 {object} u.Response{Success bool, Message string, Errors any}
// @Router /auth/login [post]
func LoginHandler(ctx *gin.Context) {
	var req m.LoginRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})
		return
	}

	if req.Email == "" {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Email is required",
		})
		return
	}

	if req.Password == "" {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Password is required",
		})
		return
	}


	if !m.MatchUserInDatabase(req.Email, req.Password) {
		ctx.JSON(http.StatusUnauthorized, u.Response{
			Success: false,
			Message: "Unauthorized. Make sure your email and password is correct",
		})
		return
	}

	token, err := u.GenerateToken(req.ID, req.Email)
	if err != nil {
		fmt.Println("LoginHandler error generate token:", err)
	}

	ctx.JSON(http.StatusOK, u.Response{
		Success: true,
		Message: "Login Success",
		Token:   token,
	})
}

// @summary Handle forgot password
// @Description Forgot password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body m.ForgotPasswordRequest true "request forgot password"
// @Success 200 {object} u.Response{Success bool, Message string, Errors any}
// @Failure 400 {object} u.Response{Success bool, Message string, Errors any}
// @Failure 404 {object} u.Response{Success bool, Message string, Errors any}
// @Router /auth/forgot-password [post]
func ForgotPasswordHandler(ctx *gin.Context) {
	var req m.ForgotPasswordRequest
	
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"message": "Internal server error",
		})
		return
	}
	
	if req.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is required",
		})
		return
	}
	
	if !m.IsEmailExist(req.Email) {
		ctx.JSON(http.StatusNotFound, gin.H {
			"error": "Email not found",
		})
		return
	}
	
	OTP := u.GenerateOTP()
	OTPReq := m.OTPRequest{
		Email: req.Email,
		OTP: OTP,
	}
	redisClient := config.RedisConnect()
	encoded, err := json.Marshal(OTPReq)
			if err != nil {
				fmt.Println("failed to marshal json:", err)
			}

	redisClient.Set(context.Background(), "users", string(encoded), time.Duration(5)*time.Minute)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Verification code sent successfully",
		"verification_code": OTP,
	})
}

func ChangePasswordHandler(ctx *gin.Context) {
	var req m.ChangePasswordRequest
	fmt.Println("req:", req)
	ctx.ShouldBind(&req)
	
	if req.NewPassword == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "New password are required",
		})
		return
	}
	
	if req.ConfirmNewPassword == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Confirm password are required",
		})
		return
	}
	
	if len(req.NewPassword) < 8 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Password must be at least 8 characters",
		})
		return
	}
	
	if req.NewPassword != req.ConfirmNewPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Password does not match!",
		})
		return
	}
	
	if !m.IsEmailExist(req.Email) {
		ctx.JSON(http.StatusNotFound, gin.H {
			"error": "Email not found",
		})
		return
	}
	
	fmt.Println("req email controller", req.Email, req.NewPassword)
	
	m.UpdateUserPassword(req.Email, req.NewPassword)
	
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "Password reset successfully",
	// 	"user": u.FindUserByEmail(req.Email),
	// })
	
}