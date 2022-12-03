package main

import (
	"fmt"
	"net/http"
)

var count int

func main()  {
	fmt.Println("サーバーを開始します。")

	http.HandleFunc("/", handler)
	http.HandleFunc("/get", countHandler)
	
	http.ListenAndServe(":8080", nil)
	fmt.Println("finish")
}

func handler(w http.ResponseWriter, r *http.Request) {
	count ++
	fmt.Fprintln(w, count)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, count)
}