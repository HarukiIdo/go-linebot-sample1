package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HarukiIdo/go-linebot-sample1/handler"
)

func main() {

	// ハンドラの登録
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/callback", handler.LineHandler)

	fmt.Println("http://localhost:8080で起動中...")

	// HTTPサーバを起動
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
