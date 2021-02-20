package main

import (
	"go_programing/chapter7/server/handler"
	"log"
	"net/http"
)

var db = handler.Database{"shoes": 25.0, "T-shit": 20.0, "paints": 99.9}

func main() {
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/query", db.QueryPrice)
	http.HandleFunc("/delete", db.DeleteItem)
	http.HandleFunc("/add", db.AddItem)
	http.HandleFunc("/update", db.UpdateItem)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}