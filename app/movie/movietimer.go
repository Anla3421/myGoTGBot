package movie

import (
	"fmt"
	"strconv"
	"time"
)

type (
	MovieData struct {
		List []string
		Page int
	}

	MovieList struct {
		RawData      []string
		ListMap      map[string]*MovieData
		Page         int
		intChan      chan int
		sliceChan    chan []string
		reqPageChan  chan string
		resPageChan  chan []string
		AddMovieChan chan *MovieData
	}
)

// 儲存整理好的資料 mainService
var mainService *MovieList

// 實體化 MovieList
func newMovielistTask() *MovieList {
	return &MovieList{
		ListMap:      make(map[string]*MovieData),
		intChan:      make(chan int, 2),
		sliceChan:    make(chan []string),
		reqPageChan:  make(chan string),
		resPageChan:  make(chan []string),
		AddMovieChan: make(chan *MovieData, 10),
	}
}

// MovieTimer:電影資料更新計時器
func MovieTimer() {
	task := newMovielistTask()
	mainService = task
	updateTime := 3600
	tickerFirst := time.NewTimer(time.Second)
	tickerUpdate := time.Duration(updateTime) * time.Second
	ticker := time.NewTicker(tickerUpdate)

	for {
		select {
		// 首次執行
		case <-tickerFirst.C:
			task.Page = 0
			task.intChan <- updateTime
		// 定時執行
		case <-ticker.C:
			task.Page = 0
			task.intChan <- updateTime
		// 執行更新
		case temp := <-task.intChan:
			fmt.Printf("設定時間%v秒已到，開始更新\n", temp)
			task.MoviespiderNew(task.sliceChan)
		// 整理資料
		case outputData := <-task.AddMovieChan:
			task.ListMap[strconv.Itoa(outputData.Page)] = outputData
		// 回傳req要求的資料
		case reqPage := <-task.reqPageChan:
			data := task.OutputData(reqPage)
			task.resPageChan <- data
		}

	}
}

// OutputData:整理回傳req要求的某頁資料
func (task *MovieList) OutputData(reqPage string) []string {
	if mainService == nil {
		return []string{"沒有這麼多電影可以推薦呢~"}
	}
	if data, exist := mainService.ListMap[reqPage]; exist {
		return data.List
	}
	return []string{"沒有這麼多電影可以推薦呢~~~"}
}

// GetMoviePage:要求某頁資料
func GetMoviePage(page string) []string {
	mainService.reqPageChan <- page
	data := <-mainService.resPageChan

	return data
}
