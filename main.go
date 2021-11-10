package main

import (
	"server/api"
	"server/api/controller/bot"
	"server/api/controller/movie"
	"server/service/nlscSpider/cache"
	"server/service/nlscSpider/checker"
	_ "server/service/nlscSpider/scraper"
)

//---------------------------------------------------------------------------
func main() {

	cacheServer := cache.NewCacheServer()
	go cacheServer.Run()

	checker.StartChecker()

	go api.Api()
	go bot.Bot()

	// movie.Moviespider() // DB & sleep
	movie.MovieTimer() // map & channel

}
