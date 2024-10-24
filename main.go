package main

import (
	"server/api"
	"server/api/controller/bot"
	"server/api/controller/movie"
	"server/service/nlscSpider/cache"
	_ "server/service/nlscSpider/scraper"
)

// ---------------------------------------------------------------------------
func main() {

	cacheServer := cache.NewCacheServer()
	go cacheServer.Run()

	// checker.StartChecker() // 天氣爬蟲啟用會有問題，先關閉

	go api.Api()
	go bot.Bot()

	// movie.Moviespider() // DB & sleep
	movie.MovieTimer() // map & channel

}
