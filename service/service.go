package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/HarukiIdo/go-linebot-sample1/model"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func SendResttoInfo(bot *linebot.Client, e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	replyMsg := fmt.Sprintf("緯度：%s\n軽度：%s", lat, lng)
	replyMsg = getRestoInfo(lat, lng)

	if _, err := bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage(replyMsg)).Do(); err != nil {
		log.Println(err)
	}
}

func getRestoInfo(lat string, lng string) string {
	apiKey := os.Getenv("APIKEY")
	url := fmt.Sprintf("http://webservice.recruit.co.jp/hotpepper/gourmet/v1/?format=json&key=%s&lat=%s&lng=%s", apiKey, lat, lng)

	// リクエストのBodyを取得
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data model.Response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	info := ""

	for _, shop := range data.Results.Shop {
		info += shop.Name + "\n" + shop.Address + "\n\n"
	}
	return info
}
