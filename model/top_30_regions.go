package model

import "go_test/db"

type ByRegion struct {
	Region            string
	TotalQuantitySold int64
	TotalRevenue      float64
}

func GetTop30Regions() ([]ByRegion, error) {
	query := `SELECT
    region,
    SUM(quantity) AS total_quantity_sold,
    SUM(total_price) AS total_revenue
FROM transactions
GROUP BY region
ORDER BY total_revenue DESC
LIMIT 30;`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regions []ByRegion

	for rows.Next() {
		var region ByRegion
		err := rows.Scan(&region.Region, &region.TotalQuantitySold, &region.TotalRevenue)

		if err != nil {
			return nil, err
		}
		regions = append(regions, region)

	}
	return regions, nil
}
