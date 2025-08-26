package routes

import (
	"github.com/gin-gonic/gin"
	"go_test/model"
	"go_test/utils"
	"net/http"
)

func getMonthlyRevenue(ctx *gin.Context) {
	result, err := model.GetMonthlyRevenue()

	// mapping the data

	months, soldQuantity, totalRevenue := utils.SeparateMonthlyData(result)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get monthly revenue"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": gin.H{
		"months":       months,
		"soldQuantity": soldQuantity,
		"totalRevenue": totalRevenue,
	},
	})
}
