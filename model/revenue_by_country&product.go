package model

import (
	"go_test/db"
)

type RevenueByCountryAndProduct struct {
	Country          string
	ProductName      string
	TotalRevenue     float64
	TransactionCount int
}

func GetRevenueByCountryAndProduct() ([]RevenueByCountryAndProduct, error) {
	query := `SELECT country, product_name, SUM(total_price) AS total_revenue, COUNT(*) AS transaction_count
            FROM transactions
            GROUP BY country, product_name
            ORDER BY total_revenue DESC`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Results []RevenueByCountryAndProduct

	for rows.Next() {
		var res RevenueByCountryAndProduct
		err := rows.Scan(&res.Country, &res.ProductName, &res.TotalRevenue, &res.TransactionCount)

		if err != nil {
			return nil, err
		}
		Results = append(Results, res)
	}
	return Results, nil

}
