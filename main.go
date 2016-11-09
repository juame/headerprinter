package main

import (
	"encoding/json"
	"net/http"
)

func responder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(r.Header)
	w.Write(j)
}

func main() {
	http.HandleFunc("/", responder)
	http.ListenAndServe(":8080", nil)
}
