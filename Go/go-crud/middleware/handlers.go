package middleware

import (
	"database/sql"
	"fmt"
	"go-psql-crud/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"encoding/json"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	er := godotenv.Load(".env")

	var (
		// host     = "localhost"
		// port     = 5432
		// user     = "postgres"
		// password = "abc123"
		// dbname   = "stocksdb"
		host     = os.Getenv("DB_HOST")
		port, _  = strconv.Atoi(os.Getenv("DB_PORT"))
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	if er != nil {
		log.Fatal("Error loading .env")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	// db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("sucessfully connected to postgres")

	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("unable to decode %v", err)
	}
	insertID := insertStock(stock)

	res := response{
		ID:      insertID,
		Message: "stock created",
	}
	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string %v", err)
	}
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatalf("Unable to get stovk %v", err)
	}
	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()
	if err != nil {
		log.Fatalf("Unable to get %v", err)
	}
	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string %v", err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode %v", err)
	}
	updatedRows := updateStock(int64(id), stock)
	msg := fmt.Sprintf("Stock updated %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {

		log.Fatalf("Unable to convert %v", err)
	}
	deletedRows := deleteStock(int64(id))
	msg := fmt.Sprintf("stock deleted %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) int64 {

	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO stocks(name,price,company) VALUES($1,$2,$3) RETURNING stockid`
	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execut %v", err)
	}
	fmt.Printf("inserted single record %v", id)
	return id
}

func getStock(id int64) (models.Stock, error) {
	db := createConnection()
	defer db.Close()
	var stock models.Stock
	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to Scan %v", err)
	}
	return stock, err

}

func getAllStocks() ([]models.Stock, error) {

	db := createConnection()
	defer db.Close()

	var stocks []models.Stock
	sqlStatement := `SELECT * FROM stocks`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Unabl eto scan %v", err)
		}
		stocks = append(stocks, stock)

	}
	return stocks, err
}

func updateStock(id int64, stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `UPDATE stocks SET name=$2,price=$3,company=$4 WHERE stockid=$1`

	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatalf("Unable to execute %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows %v", err)
	}
	fmt.Printf("total rows %v", rowsAffected)
	return rowsAffected

}

func deleteStock(id int64) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows %v", err)
	}
	fmt.Printf("total rows %v", rowsAffected)
	return rowsAffected
}