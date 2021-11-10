package callback

import (
	"encoding/json"
	"fmt"
	"server/service/nlscSpider/cache"
	"server/service/nlscSpider/lib"

	"server/service/mylib/errorCode"
	"server/service/mylib/selfTime"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	PerWeatherInfo = "p"
)

type GetPerWeatherInfoReq struct {
	LocationName string `json:"l"`
}

func GetPerWeatherInfo(message *tgbotapi.Update, req []byte) (code int, res tgbotapi.MessageConfig, err error) {
	request := &GetPerWeatherInfoReq{}
	err = json.Unmarshal(req, request)
	fmt.Println(req)
	fmt.Println(request)
	fmt.Println(request.LocationName)
	if err != nil {
		code = errorCode.DecodeJsonError
		return
	}

	cache.Server.GetWeatherDataReq <- request.LocationName
	weatherInfo := <-cache.Server.GetWeatherDataRes

	if weatherInfo == nil {
		code = errorCode.DBNoData
		return
	}

	code, content, err := lib.FormatWeatherData(weatherInfo)
	if err != nil {
		return
	}

	content += "最後更新時間" + weatherInfo.UpdateTime.Format(selfTime.TimeLayout) + "\n"

	res = lib.NewResponseMs(message.CallbackQuery.Message.Chat.ID, content)
	return
}
