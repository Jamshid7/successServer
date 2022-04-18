package botops

import (
	"time"

	"github.com/gin-gonic/gin"
	botAPI "github.com/go-telegram-bot-api/telegram-bot-api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	mybot *botAPI.BotAPI
	upd   botAPI.UpdateConfig
	msgs  Messages
)

func Start() (error, string) {
	mybot, _ = botAPI.NewBotAPI("2059285433:AAEqN_qSF3G15T5FydC-7CWS6hZW2vCtJaE")

	upd = botAPI.NewUpdate(0)
	upd.Timeout = 60

	return nil, "I started my job!"
}

func InitMessages(msg *Messages) {
	msgs = *msg
}

func SendMessage() {
	getMsg := getByPriority(&msgs)
	mybot.Send(botAPI.NewMessageToChannel("-1001721270619", getMsg))
}

// @Title Telegram API Server
// @Version 1.0.0
// @Summary Send message to channel
// @Description API to send simple message
// @Router /send [GET]
// @Success 200
// @Failure 422
// @Failure 500
func SendToChannel(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delivered",
	})
	time.AfterFunc(time.Second*5, SendMessage)
}

func Router() {
	r := gin.Default()

	r.GET("/send", SendToChannel)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

func getByPriority(msg *Messages) string {
	for _, val := range msg.Msgs {
		if val.Priority == "high" {
			tempTxt := val.Text
			msg.Delete(val)
			return tempTxt
		}
	}
	for _, val := range msg.Msgs {
		if val.Priority == "medium" {
			tempTxt := val.Text
			msg.Delete(val)
			return tempTxt
		}
	}
	for _, val := range msg.Msgs {
		if val.Priority == "low" {
			tempTxt := val.Text
			msg.Delete(val)
			return tempTxt
		}
	}
	return ""
}
