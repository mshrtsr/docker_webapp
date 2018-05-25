package main

import (
	"bytes"
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request){
	
	ContentType := "application/json"

	if r.Header.Get("Content-Type") != ContentType {
		w.WriteHeader(404)
		return	
	}

	if r.Method != "GET" {
		w.WriteHeader(404)
		return
	}

	type Response struct {
		Message string `json:"message"`
	}

	response := Response {
		Message: "Hello World!!",
	}

	jsonData, _ := json.Marshal(response)
	jsonFormedData := new(bytes.Buffer)
	json.Indent(jsonFormedData, jsonData, "", "    ")	

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, jsonFormedData.String())

}

func main(){
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

