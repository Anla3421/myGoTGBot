package movie

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"server/domain/repository/dao"
	"strings"
)

// 建立請求
var client = &http.Client{}

func fetch(url string) string {
	fmt.Println("Fetch Url", url)

	// 建立HTTP客戶端
	req, _ := http.NewRequest("GET", url, nil)
	// 發出請求
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	resp, err := client.Do(req)
	// 錯誤回報
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	// 檢查網頁狀態(code)如果有錯誤就回報
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	// 讀取HTTP響應正文
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	// 錯誤回報
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)
}
func parseUrls(url string, i int) {
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
		dao.Moviesql(ID, IdResql, TitleResql)
		ID = ID + 1
	}
}
