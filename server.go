package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request){
	
	type Response struct {
		Message string `json:"message"`
	}

	response := Response {
		Message: "\nHello World!!\n",
	}

	json, _ := json.Marshal(response)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(json)

}

func main(){
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

