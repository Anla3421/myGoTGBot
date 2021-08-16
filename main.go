package main

import (
	"server/api"
	"server/api/controller/bot"
	"server/api/controller/movie"
)

//---------------------------------------------------------------------------
func main() {
	go api.Api()
	go movie.Moviespider()
	bot.Bot()
}
