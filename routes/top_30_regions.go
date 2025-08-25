package routes

import (
	"github.com/gin-gonic/gin"
	"go_test/model"
	"net/http"
)

func getTop30Regions(ctx *gin.Context) {
	result, err := model.GetTop30Regions()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get top 30 regions",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
