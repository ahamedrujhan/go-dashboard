package routes

import (
	"github.com/gin-gonic/gin"
	"go_test/model"
	"net/http"
)

func getTop20Products(ctx *gin.Context) {
	result, err := model.GetTop20Products()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get top20 products"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
