package checker

import "server/infrastructure/service/nlscSpider/config"

func StartChecker() {

	newWeatherChcker := NewWeatherChecker(config.ServerConfig.WeatherChecker)
	go newWeatherChcker.Run()

}
