package bot

import (
	"fmt"
	"log"
	"server/model/dao"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Bot() {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := BotConn.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	fmt.Print("Success connected, bot online.")

	for update := range updates {
		if update.CallbackQuery != nil {

			BotConn.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			switch update.CallbackQuery.Data {
			//第二層面板喚醒
			case "天氣查詢":
				msg.ReplyMarkup = WeatherKB
			case "豆瓣網電影精選":
				msg.ReplyMarkup = MovieKB
			//天氣選項回應
			case "5": //我全都要
				msg.Text = ""
				for ID := 1; ID < 5; ID++ {
					result := dao.Weathersql(strconv.Itoa(ID)).Text + "\n"
					msg.Text += result
				}
			default: //天氣單項1~4
				msg.Text = ""
				result := dao.Weathersql(update.CallbackQuery.Data).Text
				if result != "" {
					msg.Text += result
				}

			}
			//電影選項回應
			if update.CallbackQuery.Message.Text == "豆瓣網電影精選" {
				ID, _ := strconv.Atoi(update.CallbackQuery.Data)
				msg.Text = ""
				for Rank := 25*ID + 1; Rank < 25*ID+26; Rank++ {
					result := dao.Moviesqlget(Rank).Idre + "  " + dao.Moviesqlget(Rank).Moviename + "\n"
					msg.Text += result
				}
				msg.Text = "電影推薦TOP" + strconv.Itoa(1+25*(ID)) + "~" + strconv.Itoa(25*(ID+1)) + ":\n" +
					msg.Text + "\n請使用https://movie.douban.com/subject/\n+上述電影的數字來進入該電影之介紹"
			}
			if update.CallbackQuery.Message.Text == "請選擇甜度" {
				msg.Text = update.CallbackQuery.Data
				Sugar = update.CallbackQuery.Data
				BotConn.Send(msg)
				msg.Text = "請選擇冰量"
				msg.ReplyMarkup = IceKB
			}
			if update.CallbackQuery.Message.Text == "請選擇冰量" {
				msg.Text = update.CallbackQuery.Data
				Ice := msg.Text
				BotConn.Send(msg)
				Drinkid = Drinkid + 1
				Who := update.CallbackQuery.Message.Chat.FirstName

				msg.Text = Who + "點了: " + Arguments + " " + Sugar + update.CallbackQuery.Data
				dao.Drinksql(Drinkid, Who, Arguments, Sugar, Ice)
			}
			BotConn.Send(msg)
		}
		//bot喚醒
		if update.Message != nil {
			if update.Message.IsCommand() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				switch update.Message.Command() {
				case "help":
					msg.Text = "type /drink 飲料名. or /drink-飲料名. to order drink.\n/total, /clear."
				case "drink":
					Arguments = update.Message.CommandArguments()
					if Arguments == "" {
						msg.Text = "type /drink 飲料名. or /drink-飲料名."
					} else {
						msg.Text = "飲料: " + update.Message.CommandArguments() + "已點餐"
						BotConn.Send(msg)
						msg.Text = "請選擇甜度"
						msg.ReplyMarkup = SugarKB
					}

				case "total":
					msg.Text = ""
					for ID := 1; ID < Drinkid+1; ID++ {
						result := strconv.Itoa(dao.Drinksqlget(ID).ID) + "." + dao.Drinksqlget(ID).Who + "\t" +
							dao.Drinksqlget(ID).Drink + " " + dao.Drinksqlget(ID).Sugar + dao.Drinksqlget(ID).Ice + "\n"
						msg.Text += result
					}
					if msg.Text == "" {
						msg.Text = "目前沒有人點餐喔"
					}
				case "clear":
					if update.Message.Chat.ID == OwnerID {
						dao.Drinksqltruncate()
						Drinkid = 0
						msg.Text = "table clear complete"
					} else {
						msg.Text = "您沒有權限執行這個指令"
					}

				default:
					msg.Text = "type /help"
				}
				BotConn.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

				switch update.Message.Text {
				case "hi":
					msg.ReplyMarkup = InitialKeyboard
					msg.Text = "您今天想要點什麼呢?"
				default:
					//對特定人士回復特定的問候語
					switch update.Message.Chat.ID {
					case OwnerID:
						msg.Text = "主人您好"
					default:
						msg.Text = "指令錯誤"
					}
					textToSend := "有人對你說，他是 " + strconv.Quote(update.Message.Chat.FirstName)
					textToSend += strconv.Quote(update.Message.Chat.LastName) + " 他說： " + update.Message.Text + ", ID是 "
					textToSend += strconv.FormatInt(update.Message.Chat.ID, 10)

					BotConn.Send(tgbotapi.NewMessage(OwnerID, textToSend))
				}
				BotConn.Send(msg)
			}
		}
	}
}
