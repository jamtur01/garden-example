package main

import (
	"fmt"
	"net/http"
)

func yo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Yo!")
}

func main() {
	http.HandleFunc("/yo", yo)
	fmt.Println("Server running...")

	http.ListenAndServe(":8080", nil)
}
