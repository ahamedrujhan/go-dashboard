package model

import (
	"go_test/db"
	"time"
)

type MonthlyRevenue struct {
	Month             time.Time `json:"month"`
	TotalQuantitySold int64
	TotalRevenue      float64
}

func GetMonthlyRevenue() ([]MonthlyRevenue, error) {
	query := `SELECT
    DATE_TRUNC('month', transaction_date) AS month,
    SUM(quantity) AS total_quantity_sold,
    SUM(total_price) AS total_revenue
FROM transactions
GROUP BY month
ORDER BY month DESC;`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var monthlyRevenues []MonthlyRevenue

	for rows.Next() {
		var monthlyRevenue MonthlyRevenue

		err := rows.Scan(&monthlyRevenue.Month, &monthlyRevenue.TotalQuantitySold, &monthlyRevenue.TotalRevenue)

		if err != nil {
			return nil, err
		}

		monthlyRevenues = append(monthlyRevenues, monthlyRevenue)
	}
	return monthlyRevenues, nil
}
