package main

import (
	"command/database"
	"command/migration"
	"command/product"
	"fmt"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // your React port
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	var dbconnection database.Config
	dbconnection.DBUser = "karthi"
	dbconnection.DBPassword = "Postgres"
	dbconnection.DBPort = 5432
	dbconnection.DBHost = "localhost"
	dbconnection.DBName = "postgres"

	err := database.InitDb(&dbconnection)
	if err != nil {
		log.Fatal(err)
	}

	migration.PostgresMigration(database.Db)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running...")
	})
	http.HandleFunc("/createproduct", product.CreateProduct)
	http.HandleFunc("/getproducts", product.GetProductDetails)

	fmt.Println("Server Started...")
	http.ListenAndServe(":8083", corsMiddleware(http.DefaultServeMux))
	// log.Fatal(http.ListenAndServe(":8083", corsMiddleware(http.DefaultServeMux)))
}
