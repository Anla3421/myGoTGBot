package callback

import (
	"encoding/json"
	"server/service/nlscSpider/cache"
	"server/service/nlscSpider/lib"

	"server/service/mylib/errorCode"
	"server/service/mylib/selfTime"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	PerWetherInfo = "p"
)

type GetPerWetherInfoReq struct {
	LocationName string `json:"l"`
}

func GetPerWetherInfo(message *tgbotapi.Update, req []byte) (code int, res tgbotapi.MessageConfig, err error) {
	request := &GetPerWetherInfoReq{}
	err = json.Unmarshal(req, request)
	if err != nil {
		code = errorCode.DecodeJsonError
		return
	}

	cache.Server.GetWeatherDataReq <- request.LocationName
	wetherInfo := <-cache.Server.GetWeatherDataRes

	if wetherInfo == nil {
		code = errorCode.DBNoData
		return
	}
	content := ""
	content += "地點:" + wetherInfo.LocationName + "\n"
	content += "預測時間區間:" + wetherInfo.StartTime + " - " + wetherInfo.EndTime + "\n"
	content += "天氣狀況:" + wetherInfo.Weather + "\n"
	content += "降雨機率:" + wetherInfo.ChanceOfRain + "\n"
	content += "最高溫:" + wetherInfo.MaxTemperature + "\n"
	content += "最低溫:" + wetherInfo.MinTemperature + "\n"
	content += "\n"
	content += "最後更新時間" + wetherInfo.UpdateTime.Format(selfTime.TimeLayout) + "\n"

	res = lib.NewResponseMs(message.CallbackQuery.Message.Chat.ID, content)
	return
}
