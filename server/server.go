package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	responseSender(w, profile)
}

func loggingHandler(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request recieved : %v\n", r)
		fmt.Printf("Path:%s\n", r.URL.Path)
		fmt.Println("Method :", r.Method)
		nextHandler.ServeHTTP(w, r)
	})
}

func responseSender(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(js))
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", loggingHandler(http.HandlerFunc(rootHandler)))
	fmt.Println("===============================================================================")
	fmt.Println("=========================SERVER STARTED========================================")
	fmt.Println("===============================================================================")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
