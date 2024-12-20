package cmd

import (
	"fmt"
	"server/infrastructure/service/mylib/errorCode"
	"server/infrastructure/service/mylib/logger"
	"server/infrastructure/service/mylib/selfTime"
	"server/infrastructure/service/nlscSpider/cache"
	"server/infrastructure/service/nlscSpider/callback"
	"server/infrastructure/service/nlscSpider/lib"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func GetWeatherInfo(message *tgbotapi.Update) (code int, res tgbotapi.MessageConfig, err error) {
	cache.Server.GetAllWeatherDataReq <- true
	weatherInfoList := <-cache.Server.GetAllWeatherDataRes

	content := ""
	i := 0
	for _, each := range weatherInfoList {

		i++
		tmpCode, tmpContent, tmpErr := lib.FormatWeatherData(each)
		if tmpCode != errorCode.Success {
			logger.Error(tmpErr)
			fmt.Println(tmpErr)
			return
		}
		content += tmpContent
		if i == len(weatherInfoList) {
			content += "最後更新時間" + each.UpdateTime.Format(selfTime.TimeLayout) + "\n"
		}
	}

	res = lib.NewResponseMs(message.Message.Chat.ID, content)
	return
}

func GetWeatherList(message *tgbotapi.Update) (code int, res tgbotapi.MessageConfig, err error) {

	cache.Server.GetAllWeatherDataReq <- true
	weatherInfoList := <-cache.Server.GetAllWeatherDataRes
	res = lib.NewResponseMs(message.Message.Chat.ID, "選取地區")

	list := [][]tgbotapi.InlineKeyboardButton{}

	// 這邊是將所有地區整理成選單
	rowList := tgbotapi.NewInlineKeyboardRow()

	for _, each := range weatherInfoList {
		fmt.Println(each.LocationName)
		qCode, data, _ := lib.SetCallBackReq(callback.PerWeatherInfo, &callback.GetPerWeatherInfoReq{
			LocationName: each.LocationName,
		})
		if qCode != errorCode.Success {
			return
		}

		rowList = append(rowList, tgbotapi.NewInlineKeyboardButtonData(each.LocationName, data))
		if len(rowList) == 4 {
			list = append(list, rowList)
			rowList = tgbotapi.NewInlineKeyboardRow()
		}
	}
	if len(rowList) != 0 {
		list = append(list, rowList)
	}

	res.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(list...)
	return
}
