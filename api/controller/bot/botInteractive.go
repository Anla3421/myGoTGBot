package bot

import (
	"fmt"
	"log"
	"os"
	"server/api/controller/movie"
	"server/model/dao"
	"server/service/mylib/errorCode"
	"server/service/mylib/logger"
	"server/service/mylib/selfTime"
	"server/service/myviper"
	"server/service/nlscSpider/cache"
	"server/service/nlscSpider/lib"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	TelegramSys *TelegramServer
)

type command struct {
	tgcommand tgbotapi.BotCommand
	fn        func(update *tgbotapi.Update) (int, tgbotapi.MessageConfig, error)
}

type TelegramServer struct {
	tgChannel    tgbotapi.UpdatesChannel
	mainBot      *tgbotapi.BotAPI
	commandList  map[string]*command
	sendChan     chan tgbotapi.Chattable
	callbackChan chan tgbotapi.CallbackConfig
	deletemsChan chan tgbotapi.DeleteMessageConfig
}

func Bot() {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := BotConn.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	fmt.Print("Success connected, bot online.")

	TelegramSys = &TelegramServer{
		tgChannel:    updates, //channel
		mainBot:      BotConn, //bot
		commandList:  make(map[string]*command),
		sendChan:     make(chan tgbotapi.Chattable),
		callbackChan: make(chan tgbotapi.CallbackConfig),
		deletemsChan: make(chan tgbotapi.DeleteMessageConfig),
	}

	// 新增command
	TelegramSys.AddCommandList("drink", "order drink")                          // drinkOrder
	TelegramSys.AddCommandList("total", "show total order list")                // OrderList
	TelegramSys.AddCommandList("clear", "clear order list")                     // ClearOrderList
	TelegramSys.AddCommandList("weather", "get all location weather info")      // GetWeatherInfo
	TelegramSys.AddCommandList("weather_list", "get one location weather info") // GetWeatherList

	// 設定command list
	_, _, err = TelegramSys.SetTgCommandList()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

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

				start := time.Now()
				ID, _ := strconv.Atoi(update.CallbackQuery.Data)
				msg.Text = ""
				// for DB use:
				// for Rank := 25*ID + 1; Rank < 25*ID+26; Rank++ {
				// 	result := dao.Moviesqlget(Rank).Idre + "  " + dao.Moviesqlget(Rank).Moviename + "\n"
				// 	msg.Text += result
				// }

				// for map use:
				msg.Text = strings.Join(movie.GetMoviePage(strconv.Itoa(ID)), " \n")

				msg.Text = "電影推薦TOP" + strconv.Itoa(1+25*(ID)) + "~" + strconv.Itoa(25*(ID+1)) + ":\n" +
					msg.Text + "\n請使用https://movie.douban.com/subject/\n+上述電影的數字來進入該電影之介紹"
				elapsed := time.Since(start)
				fmt.Printf("Took %s\n", elapsed)
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
			if update.CallbackQuery.Message.Text == "選取地區" {

				cache.Server.GetWeatherDataReq <- update.CallbackQuery.Data
				weatherInfo := <-cache.Server.GetWeatherDataRes

				if weatherInfo == nil {
					code := errorCode.DBNoData
					fmt.Printf("weatherInfo is nil!!!!! code:%v", code)
					return
				}

				code, content, err := lib.FormatWeatherData(weatherInfo)
				if err != nil {
					fmt.Println(code, content, err)
					return
				}

				content += "最後更新時間" + weatherInfo.UpdateTime.Format(selfTime.TimeLayout) + "\n"
				msg.Text = content
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
				case "weather":
					msg.Text = ""
					cache.Server.GetAllWeatherDataReq <- true
					weatherInfoList := <-cache.Server.GetAllWeatherDataRes
					i := 0
					for _, each := range weatherInfoList {
						i++
						tmpCode, tmpContent, tmpErr := lib.FormatWeatherData(each)
						if tmpCode != errorCode.Success {
							logger.Error(tmpErr)
							fmt.Printf("code:%v,content:%s,err:%s", tmpCode, tmpContent, tmpErr)
							return
						}
						msg.Text += tmpContent
						if i == len(weatherInfoList) {
							msg.Text += "最後更新時間" + each.UpdateTime.Format(selfTime.TimeLayout) + "\n"
						}
					}
				case "weather_list":
					msg.Text = ""
					cache.Server.GetAllWeatherDataReq <- true
					weatherInfoList := <-cache.Server.GetAllWeatherDataRes

					msg.Text = "選取地區"

					list := [][]tgbotapi.InlineKeyboardButton{}

					// 這邊是將所有地區整理成選單
					rowList := tgbotapi.NewInlineKeyboardRow()

					for _, each := range weatherInfoList {
						rowList = append(rowList, tgbotapi.NewInlineKeyboardButtonData(each.LocationName, each.LocationName))
						if len(rowList) == 4 {
							list = append(list, rowList)
							rowList = tgbotapi.NewInlineKeyboardRow()
						}
					}
					if len(rowList) != 0 {
						list = append(list, rowList)
					}
					msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(list...)

				case "clear":
					if update.Message.Chat.ID == myviper.New().GetInt64("OwnerID") {
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
					case myviper.New().GetInt64("OwnerID"):
						msg.Text = "主人您好"
					default:
						msg.Text = "指令錯誤"
					}
					textToSend := "有人對你說，他是 " + strconv.Quote(update.Message.Chat.FirstName)
					textToSend += strconv.Quote(update.Message.Chat.LastName) + " 他說： " + update.Message.Text + ", ID是 "
					textToSend += strconv.FormatInt(update.Message.Chat.ID, 10)

					BotConn.Send(tgbotapi.NewMessage(myviper.New().GetInt64("OwnerID"), textToSend))
				}
				BotConn.Send(msg)
			}
		}
	}
}

// 將可用的指令 統一整理
func (server *TelegramServer) AddCommandList(act string, des string) {
	server.commandList[act] = &command{
		tgcommand: tgbotapi.BotCommand{
			Command:     act,
			Description: des,
		},
	}
}

// 設定機器人的menu list
func (server *TelegramServer) SetTgCommandList() (code int, data interface{}, err error) {
	tgCommandList := []tgbotapi.BotCommand{}
	for _, command := range server.commandList {
		if command.tgcommand.Command != "start" {
			tgCommandList = append(tgCommandList, command.tgcommand)
		}
	}
	err = server.mainBot.SetMyCommands(tgCommandList)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
