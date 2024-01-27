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
	"github.com/src/todoService/http/todo"
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

	fmt.Println("Connected to the database!")

	todoStore := todoStr.NewTodoStore(db)
	todoService := todoSvc.NewService(todoStore)
	todoHttp := todo.NewHttpHandler(todoService)
	http.HandleFunc("/api", todoHttp.TodoHandler)

	// Starting the HTTP server on port 5454.
	err = http.ListenAndServe(":5454", nil)
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
			envVars[row[0]] = row[1]
		}
	}

	return envVars, nil
}
