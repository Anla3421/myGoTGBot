package checker

import "server/service/nlscSpider/config"

func StartChecker() {

	newWeatherChcker := NewWeatherChecker(config.ServerConfig.WeatherChecker)
	go newWeatherChcker.Run()

}
