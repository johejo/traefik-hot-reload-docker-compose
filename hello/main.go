package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", debugHandler)
	log.Println(http.ListenAndServe("0.0.0.0:9001", mux))
}

func debugHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	req := struct {
		Method  string              `json:"method"`
		URL     string              `json:"url"`
		Headers map[string][]string `json:"headers"`
	}{
		Method:  r.Method,
		URL:     r.URL.String(),
		Headers: r.Header,
	}
	b, err := json.Marshal(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}
