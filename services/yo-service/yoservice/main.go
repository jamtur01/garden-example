package main

import (
	"fmt"
        "io/ioutil"
        "log"
	"net/http"
)

func yo(w http.ResponseWriter, r *http.Request) {
        res, err := http.Get("http://mate-service/mate")
	if err != nil {
		log.Fatal(err)
	}
	mate, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, "Yo", " ", string(mate))
}

func main() {
	http.HandleFunc("/yo", yo)

	fmt.Println("Server running...")
        log.Fatal(http.ListenAndServe(":8080", nil))
}
