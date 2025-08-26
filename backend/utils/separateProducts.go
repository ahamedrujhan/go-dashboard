package utils

import "go_test/model"

func SeparateProducts(data []model.ByProducts) ([]string, []int64, []int64) {
	var products []string
	var quantitySold []int64
	var stockQuantity []int64

	for _, d := range data {
		products = append(products, d.ProductName)
		quantitySold = append(quantitySold, d.TotalQuantitySold)
		stockQuantity = append(stockQuantity, d.StockQuantity)
	}
	return products, quantitySold, stockQuantity
}
