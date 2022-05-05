package service

import (
	"fmt"
	"log"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func SendResttoInfo(bot *linebot.Client, e *linebot.Event) {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	replyMsg := fmt.Sprintf("緯度：%s\n軽度：%s\n", lat, lng)

	if _, err := bot.ReplyMessage(e.ReplyToken, linebot.NewTextMessage(replyMsg)).Do(); err != nil {
		log.Println(err)
	}
}
