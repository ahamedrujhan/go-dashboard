package routes

import (
	"github.com/gin-gonic/gin"
	"go_test/model"
	"net/http"
)

func getRevenueByCountryAndProduct(ctx *gin.Context) {
	result, err := model.GetRevenueByCountryAndProduct()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get revenue by country and product"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
