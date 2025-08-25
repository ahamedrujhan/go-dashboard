package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/lib/pq"
	"go_test/model"
	"os"
)

// DB Database global variable
var DB *sql.DB

func InitDB() {
	var err error

	// Initiate the db connection

	DB, err = sql.Open("postgres", "host=localhost user=root password=toor sslmode=disable dbname=go_test")

	if err != nil {
		panic(err)
	}

	// initiating the create table function
	createTables()
	initData()

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	defer DB.Close()

}

func createTables() {

	createTable := `CREATE TABLE IF NOT EXISTS transactions (
    transaction_id   VARCHAR(50) PRIMARY KEY,
    transaction_date DATE,
    user_id          VARCHAR(50),
    country          VARCHAR(100),
    region           VARCHAR(100),
    product_id       VARCHAR(50),
    product_name     VARCHAR(255),
    category         VARCHAR(100),
    price            NUMERIC(12,2),
    quantity         INT,
    total_price      NUMERIC(12,2),
    stock_quantity   INT,
    added_date       DATE
)`
	// create table
	_, err := DB.Exec(createTable)

	if err != nil {
		panic("Could not create transactions table")
	}
}

func getTableRowCount() int64 {
	var count int64
	err := DB.QueryRow("SELECT COUNT(*) FROM transactions").Scan(&count)

	if err != nil {
		panic("Could not get table count")
	}

	fmt.Println("Transaction Count is:", count)

	return count

}

func importData() {

	// open csv

	file, err := os.OpenFile("./GO_test_5m.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic("Could not open data file")
	}

	defer file.Close()

	//stream rows
	transactionChain := make(chan model.Transaction)
	go func() {
		defer close(transactionChain)
		err := gocsv.UnmarshalFile(file, &transactionChain)
		if err != nil {
			panic(err)
		}
	}()

	// Insert into batches
	batchSize := 1000
	var batch []model.Transaction

	for t := range transactionChain {
		batch = append(batch, t)
		if len(batch) >= batchSize {
			// save batch
			saveBatch(batch)

			batch = batch[:0] // reset
		}
	}

	// save remaining
	if len(batch) > 0 {
		// save batch
		saveBatch(batch)
	}

	fmt.Println("Data imported successfully...")
}

func saveBatch(batch []model.Transaction) {
	tx, err := DB.Begin()

	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare(`
INSERT INTO transactions 
(transaction_id, transaction_date, user_id, country, region, product_id, product_name, category, price, quantity, total_price, stock_quantity, added_date)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
`)
	if err != nil {
		panic(err)
	}
	for _, t := range batch {
		_, err := stmt.Exec(t.TransactionID, t.TransactionDate, t.UserID, t.Country, t.Region,
			t.ProductID, t.ProductName, t.Category, t.Price, t.Quantity,
			t.TotalPrice, t.StockQuantity, t.AddedDate)
		if err != nil {
			panic(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func copyData() {

	ctx := context.Background()
	conn, _ := pgconn.Connect(ctx, "postgres://root:toor@localhost:5432/go_test")
	defer conn.Close(ctx)

	file, err := os.Open("./GO_test_5m.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	copySQL := `COPY transactions(transaction_id,transaction_date,user_id,country,region,product_id,product_name,category,price,quantity,total_price,stock_quantity,added_date) FROM STDIN WITH CSV HEADER`

	// Stream CSV data to Postgres
	_, err = conn.CopyFrom(ctx, file, copySQL)
	if err != nil {
		fmt.Println("COPY failed:", err)
		return
	}
	fmt.Println("CSV loaded successfully!")
}

func initData() {
	count := getTableRowCount()

	if count == 0 {
		// need to import csv to db
		//importData()
		copyData()
	}
}
