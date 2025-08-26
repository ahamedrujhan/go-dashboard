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

func InitDB(conf config.Config) {
	var err error

	// Initiate the db connection

	dbstr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Username, conf.Password, conf.Db_name)

	// Initiate PgConn

	pgConnStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", conf.Username, conf.Password, conf.Host, conf.Db_name)
	ctx := context.Background()
	conn, err := pgconn.Connect(ctx, pgConnStr)

	if err != nil {
    panic(fmt.Sprintf("Failed to connect with pgconn: %v", err))
}
	

	DB, err = sql.Open("postgres", dbstr)

	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// initiating the create table function
	createTables()
	initData(conn, ctx)

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

func copyData(conn *pgconn.PgConn, ctx context.Context) {

	//dbstr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", conf.Username, conf.Password, conf.Host, conf.Db_name)

	// record the starting time
	start := time.Now()

	fmt.Println("data import started...")

	//conn, _ := pgconn.Connect(ctx, "postgres://root:toor@localhost:5432/go_test")

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

func initData(conn *pgconn.PgConn, ctx context.Context) {
	count := getTableRowCount()

	if count == 0 {
		// need to import csv to db
		//importData()
		copyData(conn, ctx)
	}
	refreshMaterializedView(conn, ctx)
	createIndex(conn, ctx)
}

func refreshMaterializedView(conn *pgconn.PgConn, ctx context.Context) {
	fmt.Println("Refreshing materialized view...")

	//  Create MV if it doesn't exist
	_, err := conn.Exec(ctx, `
		CREATE MATERIALIZED VIEW IF NOT EXISTS revenue_by_country_product AS
		SELECT country, product_name,
			   SUM(total_price) AS total_revenue,
			   COUNT(*) AS transaction_count
		FROM transactions
		GROUP BY country, product_name
		ORDER BY total_revenue DESC;
	`).ReadAll()
	if err != nil {
		fmt.Println("Failed to create MV:", err)
		return
	}

	//  Create MV if it doesn't exist
	_, err = conn.Exec(ctx, `
		CREATE MATERIALIZED VIEW IF NOT EXISTS top_products AS SELECT
    product_name,
    SUM(quantity) AS total_quantity_sold,
    MAX(stock_quantity) AS stock_quantity
FROM transactions
GROUP BY product_name;
	`).ReadAll()
	if err != nil {
		fmt.Println("Failed to create MV:", err)
		return
	}

	_, err = conn.Exec(ctx, `
		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_indexes 
				WHERE tablename = 'revenue_by_country_product' 
				AND indexname = 'idx_revenue_country_product'
			) THEN
				CREATE UNIQUE INDEX idx_revenue_country_product
				ON revenue_by_country_product (country, product_name);
			END IF;
		END
		$$;
	`).ReadAll()

	if err != nil {
		fmt.Println("Failed to create unique index:", err)
		return
	}

	// Step 3: Refresh MV concurrently
	_, err = conn.Exec(ctx, "REFRESH MATERIALIZED VIEW CONCURRENTLY revenue_by_country_product;").ReadAll()
	if err != nil {
		fmt.Println("Failed to refresh MV:", err)
		return
	}

	fmt.Println("Materialized view refreshed successfully!")
}

func createIndex(conn *pgconn.PgConn, ctx context.Context) {
	indexSQL := `
        CREATE INDEX IF NOT EXISTS idx_transactions_country_product
        ON transactions(country, product_name);
    `

	_, err := conn.Exec(ctx, indexSQL).ReadAll()
	if err != nil {
		fmt.Println("Failed to create index:", err)
		return
	}

	fmt.Println("Index created successfully (or already exists)!")
}
