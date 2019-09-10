package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	Name    string
	Hobbies []string
	Email   string
	Phone   int
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	user := User{"Alex", []string{"snowboarding", "programming"}, "alex@gmail.com", 123456}
	respondWithJSON(w, http.StatusOK, user)
}

func loggingHandler(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request recieved : %v\n", r)
		fmt.Printf("Path:%s\n", r.URL.Path)
		fmt.Println("Method :", r.Method)
		nextHandler.ServeHTTP(w, r)
	})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func connect() {
	URL, _ := os.LookupEnv("CONNECTION_URL")
	db, err := sql.Open("mysql", URL)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	log.Print("============= CONNECTED TO DB: " + URL + " =================")
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

func main() {
	mux := http.NewServeMux()
	port, _ := os.LookupEnv("PORT")
	connect()
	mux.Handle("/user", loggingHandler(http.HandlerFunc(rootHandler)))
	log.Print("===============================================================================")
	log.Print("=========================SERVER STARTED AT PORT " + port + "==========================")
	log.Print("===============================================================================")
	log.Fatal(http.ListenAndServe(port, mux))
}
