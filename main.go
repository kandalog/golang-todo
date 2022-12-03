package main

import (
	"fmt"
	"net/http"
)

var count int

func main()  {
	fmt.Println("サーバーを開始します。")

	// http.HandleFunc("/", )
	http.HandleFunc("/countup", handler)
	http.HandleFunc("/get", countHandler)
	
	http.ListenAndServe(":8080", nil)
	fmt.Println("finish")
}

func handler(w http.ResponseWriter, r *http.Request) {
	count ++
	fmt.Fprintln(w, count)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, count)
	fmt.Fprintln(w, "<html><h1>hoge</h1></html>")
}