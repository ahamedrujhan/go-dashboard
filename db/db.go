package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/lib/pq"
	"go_test/config"
	"os"
	"time"
)

// DB Database global variable
var DB *sql.DB

func InitDB() {
	var err error

	// load config
	conf := config.LoadConfig()

	// Initiate the db connection

	dbstr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Username, conf.Password, conf.Db_name)

	DB, err = sql.Open("postgres", dbstr)

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// initiating the create table function
	createTables()
	initData()

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

	fmt.Println("Transaction row Count is:", count)

	return count

}

func copyData() {

	// record the starting time
	start := time.Now()

	fmt.Println("data import started...")

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

	elapsed := time.Since(start)
	fmt.Println("Data import took %s", elapsed)
}

func initData() {
	count := getTableRowCount()

	if count == 0 {
		// need to import csv to db
		//importData()
		copyData()
	}
}
