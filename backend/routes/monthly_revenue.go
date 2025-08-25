package routes

import (
	"github.com/gin-gonic/gin"
	"go_test/model"
	"net/http"
)

func getMonthlyRevenue(ctx *gin.Context) {
	result, err := model.GetMonthlyRevenue()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get monthly revenue"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
