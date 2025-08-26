# Go Dashboard

A modern web dashboard built with Go, featuring real-time data visualization and monitoring capabilities.

## Features

- **🚀 Optimized Data Fetching**: Advanced indexing and materialized views for lightning-fast data retrieval
- **🔄 Auto Data Import**: Automatic CSV data import 
- **Real-time Dashboard**: Interactive web interface
- **REST-ful API**: Clean API endpoints for data management
- **Database Integration**: Support for PostgreSQL
- **Docker Support**: Easy deployment with Docker and Docker Compose

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: React js
- **Database**: PostgreSQL
- **Containerization**: Docker & Docker Compose

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- Go 1.19 or higher
- Docker and Docker Compose
- Git
- Node.js (if using frontend framework)

## Quick Start with Docker Compose

The easiest way to get started is using Docker Compose:

### 1. Clone the Repository

```bash
git clone https://github.com/ahamedrujhan/go-dashboard.git
cd go-dashboard
```

### 2. Copy the csv file
Copy your CSV data file to the backend directory:

```bash
# Place your CSV file in the backend folder
cp GO_test_5m.csv backend/
```

> **📝 Note**: Ensure the CSV file `GO_test_5m.csv` is placed in the `backend/` directory before starting the services.


### 3. Start with Docker Compose

open console in root of the project

```bash
# Start all services
docker-compose up -d

or 

docker compose up 

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

Once all services are running:

- **📊 Dashboard**: [http://localhost](http://localhost)
- **🔌 API**: [http://localhost:8080](http://localhost:8080)
## Initialize Database

🎉 **No manual database setup required!** 

When the backend starts, it will automatically:
- Import data from the CSV file
- Create necessary database indexes
- Set up materialized views for optimal performance


## Manual Setup

If you prefer to run the application manually:

### 1. Clone and Setup

```bash
git clone https://github.com/ahamedrujhan/go-dashboard.git
cd go-dashboard
```

### 2. Install backend Dependencies

```bash
# Go to backend directory

# modify the config.yaml file with your db credentials 

# paste the `GO_test_5m.csv` file in the backend directory 

# Install Go dependencies
go mod download

#run the backend 
go run
```

### 3. Install the frontend dependencies 

```bash
# Go to frontend/go-dashboard directory

# Install frontend dependencies (if applicable)
npm install

#run frontend 
npm run dev
```


#### PostgreSQL

```bash
# need to install the postgresql and configure the proper db credentials in backend config.yaml file
```


## API Documentation

The API provides the following endpoints:

### Dashboard
| Method | Endpoint | Description | Response |
|--------|----------|-------------|----------|
| `GET` | `/api/v1/country-product-revenue` | Get revenue data grouped by country and product | JSON array |
| `GET` | `/api/v1/top-30-regions` | Get top 30 performing regions by revenue | JSON array |
| `GET` | `/api/v1/top-20-product` | Get top 20 best-selling products | JSON array |
| `GET` | `/api/v1/monthly-revenue` | Get monthly revenue breakdown | JSON array |

### Example API Responses

#### Country Product Revenue
```bash
curl -X GET http://localhost:8080/api/v1/country-product-revenue
```

#### Top 30 Regions
```bash
curl -X GET http://localhost:8080/api/v1/top-30-regions
```

#### Top 20 Products
```bash
curl -X GET http://localhost:8080/api/v1/top-20-product
```

#### Monthly Revenue
```bash
curl -X GET http://localhost:8080/api/v1/monthly-revenue
```

> **📝 Note**: All endpoints return JSON data optimized for dashboard visualization and are designed to handle large datasets efficiently through database indexing and materialized views.


---

**Enjoy! 🚀**
