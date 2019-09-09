package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func main() {
	mux := http.NewServeMux()
	mux.Handle("/user", loggingHandler(http.HandlerFunc(rootHandler)))
	fmt.Println("===============================================================================")
	fmt.Println("=========================SERVER STARTED========================================")
	fmt.Println("===============================================================================")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
