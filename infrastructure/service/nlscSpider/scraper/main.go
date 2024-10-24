package scraper

import "server/infrastructure/service/nlscSpider/config"

func init() {
	LocationC = NewLocationScraper(config.ServerConfig.LocationConfig)
	go LocationC.Run()

	WeatherC = NewWeatherScraper(config.ServerConfig.WeatherChecker)
	go WeatherC.Run()
}
