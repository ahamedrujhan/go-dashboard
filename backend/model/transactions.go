package model

import "time"

// Transaction represents a single transaction record
type Transaction struct {
	TransactionID   string    `json:"transaction_id" csv:"transaction_id"`
	TransactionDate time.Time `json:"transaction_date" csv:"transaction_date"`
	UserID          string    `json:"user_id" csv:"user_id"`
	Country         string    `json:"country" csv:"country"`
	Region          string    `json:"region" csv:"region"`
	ProductID       string    `json:"product_id" csv:"product_id"`
	ProductName     string    `json:"product_name" csv:"product_name"`
	Category        string    `json:"category" csv:"category"`
	Price           float64   `json:"price" csv:"price"`
	Quantity        int       `json:"quantity" csv:"quantity"`
	TotalPrice      float64   `json:"total_price" csv:"total_price"`
	StockQuantity   int       `json:"stock_quantity" csv:"stock_quantity"`
	AddedDate       time.Time `json:"added_date" csv:"added_date"`
}
