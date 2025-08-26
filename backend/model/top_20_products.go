package model

import "go_test/db"

type ByProducts struct {
	ProductName       string
	TotalQuantitySold int64
	StockQuantity     int64
}

func GetTop20Products() ([]ByProducts, error) {
	//	query := `SELECT
	//    product_name,
	//    SUM(quantity) AS total_quantity_sold,
	//    MAX(stock_quantity) AS stock_quantity
	//FROM transactions
	//GROUP BY product_name
	//ORDER BY total_quantity_sold DESC
	//LIMIT 20`

	newQuery := `
SELECT *
FROM top_products
ORDER BY total_quantity_sold DESC
LIMIT 20`

	rows, err := db.DB.Query(newQuery)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []ByProducts

	for rows.Next() {
		var product ByProducts
		err := rows.Scan(&product.ProductName, &product.TotalQuantitySold, &product.StockQuantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
