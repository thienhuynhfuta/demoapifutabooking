package main

import (
	"futaapibooking/dbconnect"
	areaservice "futaapibooking/service"
	"log"
	"net/http"
)

func main() {

	dbconnect.InnitClientElasticsearch()

	http.HandleFunc("/area", func(w http.ResponseWriter, r *http.Request) {
		err := areaservice.GetAreaByKey(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":9197", nil))
}
