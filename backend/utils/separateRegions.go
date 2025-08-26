package utils

import "go_test/model"

func SeparateRegions(data []model.ByRegion) ([]string, []int64, []float64) {
	var regions []string
	var totalRevenue []float64
	var soldQuantity []int64

	for _, d := range data {
		regions = append(regions, d.Region)
		totalRevenue = append(totalRevenue, d.TotalRevenue)
		soldQuantity = append(soldQuantity, d.TotalQuantitySold)
	}
	return regions, soldQuantity, totalRevenue
}
