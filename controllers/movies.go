package controllers

import (
	m "backend-cinemax/models"
	u "backend-cinemax/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @summary Get now showing movies with limit, page, sort_by, search
// @description Get now showing movies
// @tags movies
// @accept json
// @produce json
// @param limit query string false "Limit"
// @param page query string false "Page"
// @param sort_by query string false "Sort by"
// @param search query string false "Search"
// @success 200 {object} u.Response
// @failure 400 {object} u.Response
// @router /movies/now-showing [get]
func GetNowShowingMoviesHandler(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	sortBy := ctx.DefaultQuery("sort_by", "title")
	search := ctx.DefaultQuery("search", "")
	if err != nil || limit <= 0 {
        limit = 10
    }
	if limit > 50 {
			limit = 50 // Batesin maksimal 50 item per page
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
			page = 1
	}


	movies, err := m.GetNowShowingMovies(sortBy, search, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, u.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve movies",
			Errors:  err.Error(),
		})
		return
	}

	if len(movies) == 0 {
		ctx.JSON(http.StatusNotFound, u.Response{
			Status:  http.StatusNotFound,
			Message: "No movies found",
		})
		return
	}

	ctx.JSON(http.StatusOK, u.Response{
		Status:  http.StatusOK,
		Success: true,
		Message: "Now showing movies retrieved successfully",
		Result: movies,
	})
}


// @summary Get up coming movies
// @description Get up coming movies
// @tags movies
// @accept json
// @produce json
// @success 200 {object} u.Response
// @failure 400 {object} u.Response
// @router /movies/up-coming [get]
func GetUpComingMoviesHandler(ctx *gin.Context) {
	var req m.NowShowingMoviesRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request parameters",
		})
		return
	}

	movies, err := m.GetUpComingMovies()
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
		Message: "Upcoming movies retrieved successfully",
		Result: movies,
		// Result: m.NowShowingMoviesResponse{
		// 	Page:   req.Page,
		// 	Limit:  req.Limit,
		// 	Total:  len(movies),
		// 	Movies: movies,
		// },
	})
}

// @summary Get movie detail
// @description Get movie detail
// @tags movies
// @accept json
// @produce json	
// @param id path string true "Movie ID"
// @success 200 {object} u.Response
// @failure 400 {object} u.Response
// @router /movies/{id} [get]
func GetMovieDetailHandler(ctx *gin.Context) {
	param := ctx.Param("id")
	
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request parameters",
		})
		return
	}
	fmt.Println(id)
	
	movie, err := m.GetMovieDetail(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, u.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve movie detail",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, u.Response{
		Status:  http.StatusOK,
		Success: true,
		Message: "Movie retrieved successfully",
		Result: movie,
	})
}