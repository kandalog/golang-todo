package main

import (
	"fmt"
	"net/http"
  "os"
  "log"
)

var htmlStr string

func main()  {
	fmt.Println("サーバーを開始します。")

  data, err := os.ReadFile("./frontend/index.html")
  if err != nil {
    log.Fatal(err)
  }

  htmlStr = string(data)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, htmlStr)
}

