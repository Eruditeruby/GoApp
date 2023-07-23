package main

import (
	"net/http"
	"os"
        "fmt"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Ruby hi!</h1>"))
}

func main() {
        fmt.Println("Hello Ruby2")  
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
        
}
