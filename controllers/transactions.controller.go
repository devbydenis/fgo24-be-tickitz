package controllers

import (
	"backend-cinemax/dto"
	"backend-cinemax/models"
	u "backend-cinemax/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @summary Handle create transaction
// @Description Create a new transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body dto.CreateTransactionRequest true "request create transaction"
// @Success 200 {object} utils.Response{Status int, Success bool, Message string}
// @Failure 400 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Failure 500 {object} utils.Response{Status int, Success bool, Message string, Result any}
// @Router /transactions [post]
func CreateTransactionHandler(ctx *gin.Context) {
	var req dto.CreateTransactionRequest
	token := ctx.GetHeader("Authorization")

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, u.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request parameters",
			Errors:  err.Error(),
		})
		return
	}

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, u.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	// ctx.JSON(http.StatusOK, u.Response{
	// 	Status:  http.StatusOK,
	// 	Message: "Test post success",
	// 	Result:  req,
	// })

	err = models.CreateTransaction(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, u.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create booking",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, u.Response{
		Status:  http.StatusOK,
		Message: "Transaction Success",
	})
}