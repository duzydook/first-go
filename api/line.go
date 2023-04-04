package api

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Message struct {
	Text string `json:"text" form:"text"`
}

var bot *linebot.Client
var line_client_id string

var lineChannelSecret string = os.Getenv("LINE_CHANNEL_SECRET")
var lineAccessToken string = os.Getenv("LINE_ACCESS_TOKEN")
var lineClientId string = os.Getenv("LINE_CLIENT_ID")

func Line(e *echo.Echo) {

	bot, _ = linebot.New(lineChannelSecret, lineAccessToken)

	e.POST("/line", pushMessage)
}

func pushMessage(c echo.Context) error {

	var message Message
	err := c.Bind(&message)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	bot.PushMessage(lineClientId, linebot.NewTextMessage(message.Text)).Do()

	return c.String(http.StatusOK, "sent")
}
