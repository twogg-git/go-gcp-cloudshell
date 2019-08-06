package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const VERSION = "v1.2"
const FORMAT = "2006-01-02 15:04:05"

func main() {

	log.Println("[-] Listening on... 8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Endpoints started and running...")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		current := time.Now().Local()
		fmt.Fprintf(w, current.Format(FORMAT))
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, VERSION)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
