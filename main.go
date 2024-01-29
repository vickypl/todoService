package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/src/todoService/handlers"
	"github.com/src/todoService/handlers/todo"
	todoSvc "github.com/src/todoService/service/todo"
	todoStr "github.com/src/todoService/store/todo"
)

func main() {
	fmt.Println("Server started at: 5454")

	envVars, err := getEnvVars("configs\\.local.env")

	if err != nil {
		fmt.Printf("Error in reading env, err: %v", err)
	}

	dbHost := envVars["DB_HOST"]
	dbPort := envVars["DB_PORT"]
	dbUser := envVars["DB_USER"]
	dbPassword := envVars["DB_PASSWORD"]
	dbName := envVars["DB_NAME"]

	connURL := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	// Open a database connection
	db, err := sql.Open("mysql", connURL)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	todoStore := todoStr.NewTodoStore(db)
	todoService := todoSvc.NewService(todoStore)
	todoHttp := todo.NewHttpHandler(todoService)

	router := mux.NewRouter()
	router.Handle("/api/create", handlers.JWTMiddlewareAuth(http.HandlerFunc(todoHttp.CreateHandler))).Methods("POST")
	router.Handle("/api/get", handlers.JWTMiddlewareAuth(http.HandlerFunc(todoHttp.GetHandler))).Methods("GET")
	router.Handle("/api/getbyid", handlers.JWTMiddlewareAuth(http.HandlerFunc(todoHttp.GetByIDHandler))).Methods("GET")
	router.Handle("/api/update", handlers.JWTMiddlewareAuth(http.HandlerFunc(todoHttp.UpdateHandler))).Methods("PUT")
	router.Handle("/api/delete", handlers.JWTMiddlewareAuth(http.HandlerFunc(todoHttp.DeleteHandler))).Methods("DELETE")

	// Starting the HTTP server on port 5454.
	err = http.ListenAndServe(":5454", router)
	if err != nil {
		fmt.Printf("Problem starting server, err %v", err)
	}
}

func getEnvVars(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	envVars := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "=")
		if len(row) == 2 {
			envVars[strings.TrimSpace(row[0])] = strings.TrimSpace(row[1])
		}
	}

	return envVars, nil
}
