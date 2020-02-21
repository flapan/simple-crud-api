package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("INFO: starting app")

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytesCount, err := w.Write([]byte("hello api"))
		if err != nil {
			log.Println("ERR: could not write")
		}
		log.Println("Number of bytes", bytesCount)
	})
	http.ListenAndServe(":8080", router)
}
