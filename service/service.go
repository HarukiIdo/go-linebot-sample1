package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"unicode/utf8"

	"github.com/HarukiIdo/go-linebot-sample1/model"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func SendResttoInfo(bot *linebot.Client, e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	replyMsg := getRestoInfo(lat, lng)

	res := linebot.NewTemplateMessage(
		"レストラン一覧",
		linebot.NewCarouselTemplate(replyMsg...).WithImageOptions("rectangle", "cover"),
	)

	if _, err := bot.ReplyMessage(e.ReplyToken, res).Do(); err != nil {
		log.Println(err)
	}
}

func getRestoInfo(lat string, lng string) []*linebot.CarouselColumn {
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

	var ccs []*linebot.CarouselColumn
	for _, shop := range data.Results.Shop {
		addr := shop.Address

		// ホットペッパーの住所欄が60文字以内という制限があるため
		// 61文字以上ある場合はそれ以降をカットする
		if 60 < utf8.RuneCountInString(addr) {
			addr = string([]rune(addr)[:60])
		}

		fmt.Println(shop.Photo.Mobile.URL)

		cc := linebot.NewCarouselColumn(
			shop.Photo.Mobile.URL,
			shop.Name,
			addr,
			linebot.NewURIAction("ホットペッパーを開く", shop.URLs.PC),
		).WithImageOptions("#225588")
		ccs = append(ccs, cc)
	}
	return ccs
}
