package controllers

import (
	"backend-cinemax/dto"
	"backend-cinemax/models"
	"backend-cinemax/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfileHandler(ctx *gin.Context) {
	email := ctx.Param("email")

	if email == "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Email is required",
		})
		return
	}

	user, err := models.GetUserByEmail(email)
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
		Result:    user,
	})
}
