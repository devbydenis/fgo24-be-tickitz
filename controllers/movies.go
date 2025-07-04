package controllers

import (
	"backend-cinemax/models"
	m "backend-cinemax/models"
	u "backend-cinemax/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNowShowingMoviesHandler(ctx *gin.Context) {

	var req m.NowShowingMoviesRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request parameters",
		})
		return
	}

	movies, err := models.NowShowingMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, u.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve movies",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, u.Response{
		Status:  http.StatusOK,
		Success: true,
		Message: "Now showing movies retrieved successfully",
		Result: m.NowShowingMoviesResponse{
			Page:   req.Page,
			Limit:  req.Limit,
			Total:  len(movies),
			Movies: movies,
		},
	})
}