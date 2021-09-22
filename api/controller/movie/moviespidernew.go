package movie

import (
	"fmt"
	"regexp"
	"strings"

	"strconv"
	"time"
)

func (task *MovieList) parseUrlsToMap(url string, i int) (MovieData MovieData) {
	body := fetch(url)
	body = strings.Replace(body, "\n", "", -1)
	rp := regexp.MustCompile(`<div class="hd">(.*?)</div>`)
	titleRe := regexp.MustCompile(`<span class="title">(.*?)</span>`)
	idRe := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(\d+)/"`)
	items := rp.FindAllStringSubmatch(body, -1)
	ID := 25*i + 1
	for _, item := range items {
		//列印爬蟲的結果，debug時再用
		//fmt.Println(idRe.FindStringSubmatch(item[1])[1],titleRe.FindStringSubmatch(item[1])[1])
		IdResql := idRe.FindStringSubmatch(item[1])[1]
		TitleResql := titleRe.FindStringSubmatch(item[1])[1]
		MovieData.List = append(MovieData.List, strconv.Itoa(ID)+" "+IdResql+" "+TitleResql)
		ID = ID + 1
	}
	return MovieData
}

func (task *MovieList) MoviespiderNew(sliceChan chan []string) {

	start := time.Now()
	for i := 0; i < 2; i++ {
		//一頁有25個電影，共10頁
		// sliceChan <- task.parseUrlsToMap("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), i)
		// MovieListMap[strconv.Itoa(i)] = MovieListRawData

		// use map
		MovieData := task.parseUrlsToMap("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), i)
		MovieData.Page = i
		task.AddMovieChan <- &MovieData

	}
	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
	fmt.Println("Moivelist update complete!")
}
