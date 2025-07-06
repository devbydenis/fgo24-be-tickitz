package controllers

import (
	"backend-cinemax/dto"
	m "backend-cinemax/models"
	"backend-cinemax/utils"
	"fmt"
	"strconv"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @summary Handle create movie with all relations
// @Description Create a new movie with all relations (genres, casts, directors)
// @Tags admin
// @Accept json
// @Produce json
// @Param movie body dto.MoviesRequest true "request create movie"
// @Success 200 {object} utils.Response{Status int, Success bool, Message string, Result dto.MoviesRequest}
// @Failure 400 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Failure 500 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Router /admin [post]
func CreateMovieAdminHandler(ctx *gin.Context) {
	var req dto.MoviesRequest

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

// @summary Handle list all movie admin
// @Description List all movie admin
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response{Status int, Success bool, Message string, Result []m.Admin}
// @Failure 500 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Router /admin/list [get]
func ListMovieAdminHandler(ctx *gin.Context) {
	// get all movie admin
	movies, err := m.GetAllMovieAdmin()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: "failed to get all movie admins",
			Result:  nil,
		})
		return
	}

	// SUCCESS
	ctx.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Success: true,
		Message: "success to get all movie",
		Total:  int64(len(movies)),
		Result:  movies,
	})
}

// @summary Handle update movie with all relations
// @Description Update a movie with all relations (genres, casts, directors)
// @Tags admin
// @Accept json
// @Produce json
// @Param movie body dto.MoviesRequest true "request update movie"	
// @Success 200 {object} utils.Response{Status int, Success bool, Message string, Result dto.MoviesRequest}	
// @Failure 400 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Failure 500 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Router /admin/update [patch]
func UpdateMovieAdminHandler(ctx *gin.Context) {
	var req dto.MoviesRequest

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

	// // check required fields
	// fieldError := utils.CheckFieldValues(req)
	// if fieldError != "" {
	// 	ctx.JSON(http.StatusBadRequest, utils.Response{
	// 		Status:  http.StatusBadRequest,
	// 		Success: false,
	// 		Message: fieldError,
	// 		Result:  nil,
	// 	})
	// 	return
	// }

	// update movie with all relations
	err = m.UpdateMovieAdmin(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: "failed to update movie with all relations",
			Result:  nil,
		})
		return
	}

	// SUCCESS
	ctx.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Success: true,
		Message: "success to update movie",
		Result:  req,
	})
}

// @summary Handle delete movie admin
// @Description Delete a movie admin by ID
// @Tags admin
// @Accept json	
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Failure 400 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Failure 500 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Router /admin/delete/{id} [delete]
func DeleteMovieAdminHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Status:  http.StatusBadRequest,
			Success: false,
			Message: "failed to convert id to int",
		})
		return
	}

	// delete movie admin by ID
	err = m.DeleteMovieAdmin(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Success: false,
			Message: "failed to delete movie admin",
		})
		return
	}

	// SUCCESS
	ctx.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Success: true,
		Message: fmt.Sprintf("success to delete movie with ID %d", id),
	})
}