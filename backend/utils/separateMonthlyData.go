package utils

import (
	"go_test/model"
)

func SeparateMonthlyData(data []model.MonthlyRevenue) ([]string, []int64, []float64) {
	var months []string
	var quantitySold []int64
	var revenue []float64
	for _, d := range data {
		months = append(months, d.Month.Format("2006-01"))
		quantitySold = append(quantitySold, d.TotalQuantitySold)
		revenue = append(revenue, d.TotalRevenue)
	}

	return months, quantitySold, revenue
}
