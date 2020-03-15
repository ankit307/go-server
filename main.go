package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Index")
}

func main() {
	fmt.Println("Server Listening at 3000")
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":3000", nil)
	fmt.Println(err)
}
