package main

import (
	"fmt"
	"log"
	"net/http"
)

func newBuildHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "ACK")
}

func main() {
	http.HandleFunc("/newBuild", newBuildHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}