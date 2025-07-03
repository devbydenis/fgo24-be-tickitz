package controllers

import (
	"backend-cinemax/models"
	m "backend-cinemax/models"
	u "backend-cinemax/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNowShowingMoviesHandler(ctx *gin.Context) {
	// This function will handle the retrieval of now showing movies.
	// It will use the NowShowingMoviesRequest and NowShowingMoviesResponse structs
	// to manage the request and response data.
	// The implementation will involve querying the database for movies that are currently showing,
	// applying any filters or sorting specified in the request, and returning the results in the response.
	// The function will also handle pagination, allowing clients to specify the number of results per page
	// and the page number they want to retrieve.
	// The response will include the list of movies, the total number of movies available,
	// the current page, and the limit of results per page.
	// The function will be implemented in the future.
	// For now, it will be a placeholder function.

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