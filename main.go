package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/HarukiIdo/go-linebot-sample1/handler"
)

func main() {

	// ハンドラの登録
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/callback", handler.LineHandler)

	fmt.Println("http://localhost:8080で起動中...")

	port := os.Getenv("PORT")
	addr := ":" + port

	// HTTPサーバを起動
	http.ListenAndServe(addr, nil)
}
