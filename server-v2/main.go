package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Payload string `json:"payload"`
	}{
		Payload: "halo bang versi baru",
	}
	byteData, err := json.Marshal(payload)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteData)
}

func main() {
	log.Println("server v2 starting")
	http.HandleFunc("/hello", HelloHandler)
	if err := http.ListenAndServe("0.0.0.0:9090", nil); err != nil {
		log.Fatal(err.Error())
		return
	}
}
