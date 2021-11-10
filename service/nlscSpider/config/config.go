package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"server/service/myviper"
	"time"
)

var (
	Group             int64 = myviper.New().GetInt64("GroupID")
	Jared             int64 = myviper.New().GetInt64("OwnerID")
	RDMember                = []int64{Jared}
	AllowChatId             = []int64{Group, Jared}
	TaiwanLocation, _       = time.LoadLocation("Asia/Taipei")
	Power                   = true
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Account  string `json:"account"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	TimeZone string `json:"timezone"`
}

type LocationConfig struct {
	Url string `json:"url"`
}

type WeatherChecker struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

type Config struct {
	LocationConfig *LocationConfig `json:"location_config"`
	WeatherChecker *WeatherChecker `json:"weather_checker"`
}

// 設定檔
var ServerConfig = &Config{
	LocationConfig: &LocationConfig{
		Url: "https://api.nlsc.gov.tw/",
	},
	WeatherChecker: &WeatherChecker{
		Token: myviper.New().GetString("WeatherToken"),
		Url:   "https://opendata.cwb.gov.tw/api",
	},
}

func init() {
	configPath := flag.String("conf", "", "set setting config path service need")
	flag.Parse()
	if *configPath != "" {
		fmt.Println("讀取設定檔")
		conf, err := ReadConfig(*configPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
			return
		}
		ServerConfig = conf
	}
	for _, member := range RDMember {
		AllowChatId = append(AllowChatId, member)
	}
}

func ReadConfig(path string) (config *Config, err error) {
	ConfigContent, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		err = readErr
		return
	}
	config = &Config{}
	if err = json.Unmarshal(ConfigContent, config); err != nil {
		return
	}
	ServerConfig = config
	return
}
