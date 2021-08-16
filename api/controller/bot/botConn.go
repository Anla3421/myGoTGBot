package bot

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	// inline button initial
	for i := 0; i < 11; i++ {
		button := tgbotapi.NewInlineKeyboardButtonData("第"+strconv.Itoa(i)+"頁", strconv.Itoa(i))
		row = append(row, button)
		if i%3 == 0 {
			total = append(total, tgbotapi.NewInlineKeyboardRow(row...))
			row = []tgbotapi.InlineKeyboardButton{}
		}
	}
	MovieKB = tgbotapi.NewInlineKeyboardMarkup(total...)

	fmt.Println("bot initial")
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	BotConn = bot
	log.Printf("Authorized on account %s", bot.Self.UserName)

}
