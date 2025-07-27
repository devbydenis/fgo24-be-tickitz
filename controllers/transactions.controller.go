package controllers

import (
	"backend-cinemax/dto"
	"backend-cinemax/models"
	"backend-cinemax/utils"
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



// @summary Handle get history
// @Description Get history
// @Tags profile
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response{Success bool, Message string}
// @Failure 400 {object} utils.Response{Success bool, Message string, Errors any}
// @Failure 500 {object} utils.Response{Success bool, Message string, Errors any}
// @Security Token
// @Router /transactions/history [get]
func GetHistoryHandler(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(string)
	token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	histories, err := models.GetHistory(userId)
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
		Message: "History retrieved successfully",
		Result: histories,
	})
}