package main

import (
	"fmt"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World")
}

func main(){
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

