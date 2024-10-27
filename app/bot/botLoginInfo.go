package bot

import (
	"server/infrastructure/service/myviper"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Input struct {
	Email   string `form:"email" json:"email" binding:"required"`
	Subject string `form:"subject" json:"subject" binding:"required"`
	Message string `form:"message" json:"message" binding:"required"`
}

// 接收 /form 網頁上的資料內容，並做格式整理後傳送給 tg 上的 ownwerID
func LoginInfo(input Input) {
	ResMsg := "bot test:\n" + "Email: " + input.Email + "\n" + "Subject: " + input.Subject + "\n" + "Message: " + input.Message + "\n"

	msg := tgbotapi.NewMessage(myviper.New().GetInt64("OwnerID"), ResMsg)
	BotConn.Send(msg)
}
