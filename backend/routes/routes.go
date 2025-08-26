package routes

import (
	"github.com/gin-gonic/gin"
	"go_test/config"
)

func RegisterRoutes(server *gin.Engine) {

	routes := server.Group("/api/v1")

	// cors config
	routes.Use(config.CorsMiddleware)

	// revenue by country and product

	// TODO : Pagination
	routes.GET("/country-product-revenue", getRevenueByCountryAndProduct)

	// Get Top 30 regions by revenue
	routes.GET("/top-30-regions", getTop30Regions)

	// Get top 20 products
	routes.GET("/top-20-products", getTop20Products)

	// Get monthly revenue
	routes.GET("/monthly-revenue", getMonthlyRevenue)

}
