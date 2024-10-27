package bot

import (
	"fmt"
	"log"
	"os"
	"server/app/botGetOrder"
	"server/app/movie"
	"server/domain/repository/dao"
	"server/infrastructure/service/mylib/errorCode"
	"server/infrastructure/service/mylib/logger"
	"server/infrastructure/service/mylib/selfTime"
	"server/infrastructure/service/myviper"
	"server/infrastructure/service/nlscSpider/cache"
	"server/infrastructure/service/nlscSpider/lib"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type command struct {
	tgcommand tgbotapi.BotCommand
	// fn        func(update *tgbotapi.Update) (int, tgbotapi.MessageConfig, error)
}

type TelegramServer struct {
	commandList  map[string]*command
	sendChan     chan tgbotapi.Chattable
	callbackChan chan tgbotapi.CallbackConfig
	deletemsChan chan tgbotapi.DeleteMessageConfig
}

func BotInit() *TelegramServer {
	return &TelegramServer{
		commandList:  make(map[string]*command),
		sendChan:     make(chan tgbotapi.Chattable),
		callbackChan: make(chan tgbotapi.CallbackConfig),
		deletemsChan: make(chan tgbotapi.DeleteMessageConfig),
	}
}

func (server *TelegramServer) Bot() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := BotConn.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Success connected, bot online.")

	// 新增 /xxx command
	server.AddCommandList()
	// 設定 command list
	_, _, err = server.SetTgCommandList()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	for update := range updates {
		// 互動式 keyborad 是否有回傳項目
		if update.CallbackQuery != nil {
			BotConn.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			switch update.CallbackQuery.Data {
			// 第二層面板喚醒
			case "天氣查詢":
				msg.ReplyMarkup = WeatherKB
			case "豆瓣網電影精選":
				msg.ReplyMarkup = MovieKB
			// 天氣選項回應
			case "5": // 我全都要
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

			switch update.CallbackQuery.Message.Text {
			// 電影選項回應
			case "豆瓣網電影精選":
				start := time.Now()
				ID, _ := strconv.Atoi(update.CallbackQuery.Data)
				msg.Text = ""
				// for DB use:
				// for Rank := 25*ID + 1; Rank < 25*ID+26; Rank++ {
				// 	result := dao.Moviesqlget(Rank).Idre + "  " + dao.Moviesqlget(Rank).Moviename + "\n"
				// 	msg.Text += result
				// }

				// for map use:
				MoiveContent := strings.Join(movie.GetMoviePage(strconv.Itoa(ID)), " \n")
				msg.Text = "電影推薦TOP" + strconv.Itoa(1+25*(ID)) + "~" + strconv.Itoa(25*(ID+1))
				msg.Text += ":\n" + MoiveContent + "\n請使用https://movie.douban.com/subject/\n+上述電影的數字來進入該電影之介紹"
				elapsed := time.Since(start)
				fmt.Printf("Took %s\n", elapsed)
			// 點飲料系統
			case "請選擇甜度":
				msg.Text = update.CallbackQuery.Data
				Sugar = update.CallbackQuery.Data
				BotConn.Send(msg)
				msg.Text = "請選擇冰量"
				msg.ReplyMarkup = IceKB
			case "請選擇冰量":
				msg.Text = update.CallbackQuery.Data
				Ice := msg.Text
				BotConn.Send(msg)
				Drinkid = Drinkid + 1
				Who := update.CallbackQuery.Message.Chat.FirstName

				msg.Text = Who + "點了: " + Arguments + " " + Sugar + update.CallbackQuery.Data
				dao.Drinksql(Drinkid, Who, Arguments, Sugar, Ice)
			case "選取地區":
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
			default:
			}

			BotConn.Send(msg)
		}
		// bot喚醒
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
					if update.Message.From.ID == myviper.New().GetInt("OwnerID") {
						dao.Drinksqltruncate()
						Drinkid = 0
						msg.Text = "table clear complete"
					} else {
						msg.Text = "您沒有權限執行這個指令"
					}
				case "orderquery":
					msg.Text = botGetOrder.OrderQuery("JWT")
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
					// 對特定人士回復特定的問候語
					switch update.Message.From.ID {
					case myviper.New().GetInt("OwnerID"):
						msg.Text = "主人您好"
					default:
						msg.Text = "指令錯誤，可以試試看hi, /help"
					}
					// 回傳訊息告知 owner 有人使用 bot
					textToSend := "有人傳訊息\n他是： " + strconv.Quote(update.Message.From.FirstName) + strconv.Quote(update.Message.From.LastName)
					textToSend += "\n他說： " + update.Message.Text
					textToSend += "\n ID是 " + strconv.FormatInt(int64(update.Message.From.ID), 10)

					BotConn.Send(tgbotapi.NewMessage(myviper.New().GetInt64("OwnerID"), textToSend))
				}
				BotConn.Send(msg)
			}
		}
	}
}

// 統一新增 /xxx 可用的 command
func (server *TelegramServer) AddCommandList() {
	needtoAdd := SlashCommandList
	for act, des := range needtoAdd {
		server.commandList[act] = &command{
			tgcommand: tgbotapi.BotCommand{
				Command:     act,
				Description: des,
			},
		}
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
	err = BotConn.SetMyCommands(tgCommandList)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
