package controllers

import (
	"backend-cinemax/dto"
	"backend-cinemax/models"
	"backend-cinemax/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfileHandler(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	fmt.Println(ctx.GetHeader("Authorization"))

	userId := ctx.MustGet("userId").(string)
	email := ctx.MustGet("email").(string)
	fmt.Println("userId di context:", userId)
	fmt.Println("email di context:", email)

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	if email == "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Email is required",
		})
		return
	}

	user, err := models.GetUserByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal Server Error",
			Errors:  err.Error(),
		})
		return
	}

	if user == (dto.GetProfileResponse{}) {
		ctx.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Profile retrieved successfully",
		Result: user,
	})
}

