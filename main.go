package main

import (
	eximgolog "eximgolog/tools"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ReadMainlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	var mylog eximgolog.LogLine
	eximgolog.InsertLogLine(mylog)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/readmainlog", ReadMainlogHandler).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
