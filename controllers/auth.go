package controllers

import (
	m "backend-cinemax/models"
	u "backend-cinemax/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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