package bot

import (
	"server/service/myviper"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Input struct {
	Email   string `form:"email" json:"email" binding:"required"`
	Subject string `form:"subject" json:"subject" binding:"required"`
	Message string `form:"message" json:"message" binding:"required"`
}

func LoginInfo(input Input) {
	ResMsg := "bot test:\n" + "Email: " + input.Email + "\n" + "Subject: " + input.Subject + "\n" + "Message: " + input.Message + "\n"

	msg := tgbotapi.NewMessage(myviper.New().GetInt64("OwnerID"), ResMsg)
	BotConn.Send(msg)
}
