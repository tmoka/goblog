package main

import (
	"log"
	"net/http"

	"github.com/tmoka/goblog/src/controller"
)

func main() {
	http.HandleFunc("/", controller.IndexHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
