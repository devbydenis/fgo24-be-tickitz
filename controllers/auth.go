package controllers

import (
	"backend-cinemax/config"
	"backend-cinemax/dto"
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
// @Param user body dto.RegisterRequest true "request create user"
// @Success 201 {object} dto.RegisterRequest
// @Failure 400 {object} u.Response{Success bool, Message string, Errors any}
// @Router /auth/register [post]
func RegisterHandler(ctx *gin.Context) {
	var req dto.RegisterRequest

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
// @Param user body dto.LoginRequest true "request login user"
// @Success 200 {object} u.Response{Success bool, Message string, Errors any}
// @Failure 400 {object} u.Response{Success bool, Message string, Errors any}
// @Failure 401 {object} u.Response{Success bool, Message string, Errors any}
// @Router /auth/login [post]
func LoginHandler(ctx *gin.Context) {
	var req dto.LoginRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Bad Request",
			Errors:  err.Error(),
		})
		return
	}

	if req.Email == "admin@gmail.com" && req.Password == "admin123"{

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

	users, err := m.MatchUserInDB(req.Email, req.Password)
	fmt.Println("users:", users)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, u.Response{
			Success: false,
			Message: "Unauthorized. Make sure your email and password is correct",
			Errors:  err.Error(),
		})
		return
	}

	// Generate token
	token, err := u.GenerateJWT(users[0].ID, users[0].Email)
	if err != nil {
		fmt.Println("LoginHandler error when generate token JWT:", err)
	}

	ctx.Header("Authorization", "Bearer " + token)

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
// @Param user body dto.ForgotPasswordRequest true "request forgot password"
// @Success 200 {object} u.Response{Success bool, Message string, Errors any}
// @Failure 400 {object} u.Response{Success bool, Message string, Errors any}
// @Failure 404 {object} u.Response{Success bool, Message string, Errors any}
// @Router /auth/forgot-password [post]
func ForgotPasswordHandler(ctx *gin.Context) {
	var req dto.ForgotPasswordRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, u.Response{
			Success: false,
			Message: "Internal Server Error",
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

	if !m.IsEmailExist(req.Email) {
		ctx.JSON(http.StatusNotFound, u.Response{
			Success: false,
			Message: "Email not found",
		})
		return
	}

	OTP := u.GenerateOTP()
	OTPReq := dto.OTPRequest{
		Email: req.Email,
		OTP:   OTP,
	}

	redisClient := config.RedisConnect()
	encoded, err := json.Marshal(OTPReq)
	if err != nil {
		fmt.Println("failed to marshal json:", err)
	}

	redisClient.Set(
		context.Background(), "users", 
		string(encoded), 
		time.Duration(5)*time.Minute,
	)

	//kirim otp via email
	// err = u.SendEmailOTP(req.Email, OTP)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, u.Response{
	// 		Success: false,
	// 		Message: "Internal Server Error",
	// 		Errors:  err.Error(),
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusOK, u.Response{
		Success: true,
		Message: "OTP sent successfully. Input before 5 minutes",
		OTP:     OTP,
	})
}

// @summary Handle change password
// @Description Change password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.ChangePasswordRequest true "request change password"
// @Success 200 {object} u.Response{Success bool, Message string, Errors any}
// @Failure 400 {object} u.Response{Success bool, Message string, Errors any}
// @Failure 404 {object} u.Response{Success bool, Message string, Errors any}
// @Router /auth/change-password [post]
func ChangePasswordHandler(ctx *gin.Context) {
	var req dto.ChangePasswordRequest
	ctx.ShouldBind(&req)

	if req.NewPassword == "" {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "New password are required",
		})
		return
	}

	if req.ConfirmNewPassword == "" {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Confirm new password are required",
		})
		return
	}

	if len(req.NewPassword) < 8 {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Password must be at least 8 characters",
		})
		return
	}

	if req.NewPassword != req.ConfirmNewPassword {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "New password and confirm new password do not match",
		})
		return
	}

	if !m.IsEmailExist(req.Email) {
		ctx.JSON(http.StatusNotFound, u.Response{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "Email not found",
		})
		return
	}

	// check validation otp 
	isOTPValid, err := m.VerifyOTP(req.OTP)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, u.Response{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
			Errors:  err.Error(),
		})
		return
	}

	if !isOTPValid {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Success: false,
			Message: "Invalid OTP",
		})
		return
	}

	// hash password before update
	HashedPassword, err := u.HashPassword(req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, u.Response{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
			Errors:  err.Error(),
		})
		return
	}

	err = m.UpdateUserPassword(req.Email, HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, u.Response{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, u.Response{
		Success: true,
		Message: "Password reset successfully",
	})
}
