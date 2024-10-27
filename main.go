package main

import (
	"server/api"
	"server/app/bot"
	"server/app/movie"
	"server/infrastructure/service/nlscSpider/cache"

	"github.com/gin-gonic/gin"
)

func main() {
	// checker.StartChecker() // 天氣爬蟲啟用會有問題，先關閉
	cacheServer := cache.NewCacheServer()
	go cacheServer.Run()
	go bot.BotInit().Bot()
	go movie.MovieTimer() // map & channel

	// init gin server
	r := gin.Default()
	api.Router(r)
	r.Run(":8000")
}
