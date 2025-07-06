package controllers

import (
	m "backend-cinemax/models"
	"backend-cinemax/utils"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @summary Handle create movie with all relations
// @Description Create a new movie with all relations (genres, casts, directors)
// @Tags admin
// @Accept json
// @Produce json
// @Param movie body m.MoviesRequest true "request create movie"
// @Success 200 {object} utils.Response{Status int, Success bool, Message string, Result m.MoviesRequest}
// @Failure 400 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Failure 500 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Router /admin [post]
func CreateAdminHandler(ctx *gin.Context) {
	var req m.MoviesRequest

	// bind request to struct
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Status:  http.StatusBadRequest,
			Success: false,
			Message: "failed to bind request",
			Result:  nil,
		})
		return
	}

	// check required fields
	fieldError := utils.CheckFieldValues(req)
	if fieldError != "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Status:  http.StatusBadRequest,
			Success: false,
			Message: fieldError,
			Result:  nil,
		})
		return
	}

	// create movie with all relations
	err = m.CreateMovieWithAllRelations(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: "failed to create movie with all relations",
			Result:  nil,
		})
		return
	}

	// SUCCESS
	ctx.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Success: true,
		Message: "success to create new movie",
		Result:  req,
	})
}
