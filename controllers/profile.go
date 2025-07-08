package controllers

import (
	"backend-cinemax/dto"
	"backend-cinemax/models"
	"backend-cinemax/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @summary Handle get profile
// @Description Get profile
// @Tags profile
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response{Success bool, Message string, Result dto.GetProfileResponse}
// @Failure 400 {object} utils.Response{Success bool, Message string, Errors any}
// @Failure 401 {object} utils.Response{Success bool, Message string, Errors any}
// @Security Token
// @Router /profile [get]
func GetProfileHandler(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	userId := ctx.MustGet("userId").(string)
	email := ctx.MustGet("email").(string)

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

// @summary Handle update profile
// @Description Update profile
// @Tags profile
// @Accept json
// @Produce json
// @Param profile body dto.UpdateProfileRequest true "request update profile"
// @Success 200 {object} utils.Response{Success bool, Message string}
// @Failure 400 {object} utils.Response{Success bool, Message string, Errors any}
// @Failure 500 {object} utils.Response{Success bool, Message string, Errors any}
// @Security Token
// @Router /profile [patch]
func UpdateProfileHandler(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(string)

	var req dto.UpdateProfileRequest	// pointer to string to make it nullable
	ctx.ShouldBind(&req)

	err := models.UpdateUser(userId, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal Server Error",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Profile updated successfully",
	})
}


// @summary Handle upload photo
// @Description Upload photo
// @Tags profile
// @Accept multipart/form-data
// @Produce json
// @Param photo formData file true "photo"
// @Success 200 {object} utils.Response{Success bool, Message string}
// @Failure 400 {object} utils.Response{Success bool, Message string, Errors any}
// @Failure 500 {object} utils.Response{Success bool, Message string, Errors any}
// @Security Token
// @Router /profile [post]
func UploadPhotoHandler(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(string)

	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Failed to upload photo",
			Errors:  err.Error(),
		})
		return
	}

	// agar supaya kalo ada yg upload dengan nama yang sama ga ke replace
	// kita lakuin hal berikut
	fileName := uuid.New().String()     // generate random string buat nama filenya
	ext := filepath.Ext(file.Filename)  // ambil extention dari metada file.Filename

	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid file format",
		})
		return
	}

	err = models.UploadPhoto(userId, fileName+ext)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal Server Error",
			Errors:  err.Error(),
		})
		return
	}

	ctx.SaveUploadedFile(file, "./uploads/"+fileName+ext) // ini relative ke file main.go (file utama)

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Photo uploaded successfully",
	})
}