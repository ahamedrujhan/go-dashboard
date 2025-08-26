package routes

import (
	"github.com/gin-gonic/gin"
	"go_test/model"
	"net/http"
	"strconv"
)

func getRevenueByCountryAndProduct(ctx *gin.Context) {

	// Get query params for pagination
	pageStr := ctx.DefaultQuery("page", "1")
	perPageStr := ctx.DefaultQuery("perPage", "50")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 50
	}

	result, err := model.GetRevenueByCountryAndProduct(page, perPage)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get revenue by country and product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":         result.Results,
		"totalRecords": result.TotalRecords,
		"totalPages":   result.TotalPages,
		"currentPage":  result.CurrentPage,
		"perPage":      result.PerPage,
	})
}
