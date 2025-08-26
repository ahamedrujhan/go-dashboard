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
type PaginatedRevenue struct {
	Results      []RevenueByCountryAndProduct
	TotalRecords int
	TotalPages   int
	CurrentPage  int
	PerPage      int
}

func GetRevenueByCountryAndProduct(page, perPage int) (*PaginatedRevenue, error) {
	//query := `SELECT country, product_name, SUM(total_price) AS total_revenue, COUNT(*) AS transaction_count
	//          FROM transactions
	//          GROUP BY country, product_name
	//          ORDER BY total_revenue DESC
	//          LIMIT $1 OFFSET $2`

	queryNew := `SELECT * FROM revenue_by_country_product
ORDER BY total_revenue DESC
LIMIT $1 OFFSET $2;`

	countQuery := `SELECT COUNT(*) 
	               FROM (SELECT country, product_name
	                     FROM transactions
	                     GROUP BY country, product_name) AS subquery`

	var totalRecords int

	err := db.DB.QueryRow(countQuery).Scan(&totalRecords)
	if err != nil {
		return nil, err
	}

	totalPages := (totalRecords + perPage - 1) / perPage
	offset := (page - 1) * perPage

	rows, err := db.DB.Query(queryNew, perPage, offset)

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

	return &PaginatedRevenue{
		Results:      Results,
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		CurrentPage:  page,
		PerPage:      perPage,
	}, nil

}
