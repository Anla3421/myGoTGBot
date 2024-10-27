package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var MovieKB tgbotapi.InlineKeyboardMarkup

var row = []tgbotapi.InlineKeyboardButton{}
var total = [][]tgbotapi.InlineKeyboardButton{}
var BotConn *tgbotapi.BotAPI

var Sugar string
var Arguments string
var Drinkid int = 0

// Inlinekeyboard Setting
var InitialKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("天氣", "天氣查詢"),
		tgbotapi.NewInlineKeyboardButtonData("電影", "豆瓣網電影精選"),
		tgbotapi.NewInlineKeyboardButtonURL("Login test page", "http://127.0.0.1:8000/form"),
	),
)
var WeatherKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("台北", "1"),
		tgbotapi.NewInlineKeyboardButtonData("台中", "2"),
		tgbotapi.NewInlineKeyboardButtonData("台南", "3"),
		tgbotapi.NewInlineKeyboardButtonData("高雄", "4"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("我全都要", "5"),
	),
)

var SugarKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("全糖", "全糖"),
		tgbotapi.NewInlineKeyboardButtonData("半糖", "半糖"),
		tgbotapi.NewInlineKeyboardButtonData("微糖", "微糖"),
		tgbotapi.NewInlineKeyboardButtonData("無糖", "無糖"),
	),
)

var IceKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("全冰", "全冰"),
		tgbotapi.NewInlineKeyboardButtonData("半冰", "半冰"),
		tgbotapi.NewInlineKeyboardButtonData("微冰", "微冰"),
		tgbotapi.NewInlineKeyboardButtonData("去冰", "去冰"),
	),
)

// /xxx commands
var SlashCommandList = map[string]string{
	"drink":        "order drink",
	"total":        "show total order list",
	"clear":        "clear order list",
	"weather":      "get all location weather info",
	"weather_list": "get one location weather info",
}
