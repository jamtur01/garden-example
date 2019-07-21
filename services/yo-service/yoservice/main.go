package main

import (
	"fmt"
        "io/ioutil"
        "log"
	"net/http"
)

func yo(w http.ResponseWriter, r *http.Request) {
        resp, err := http.Get("http://mate-service/mate")
        if err != nil {
          log.Println(err)
        }

        resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != bil {
          log.Println(err)
        }

	fmt.Fprint(w, "Yo", string(body), "!")
}

func main() {
	http.HandleFunc("/yo", yo)
	fmt.Println("Server running...")

	http.ListenAndServe(":8080", nil)
}
