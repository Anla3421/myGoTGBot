package cache

type CacheServer struct {
	//http
	SetHttpData       chan *HttpData
	GetHttpDataReq    chan string
	GetHttpDataRes    chan *HttpData
	GetAllHttpDataReq chan bool
	GetAllHttpDataRes chan map[string]*HttpData
	HttpDataList      map[string]*HttpData `json:"http_checker_status_list"` // http checker list
	// mysql
	SetMysqlData       chan *MysqlData
	GetMysqlDataReq    chan string
	GetMysqlDataRes    chan *MysqlData
	GetAllMysqlDataReq chan bool
	GetAllMysqlDataRes chan map[string]*MysqlData
	MysqlDataList      map[string]*MysqlData `json:"mysql_checker_status_list"` // mysql checker list

	// weather
	SetWeatherData       chan *WeatherData
	GetWeatherDataReq    chan string
	GetWeatherDataRes    chan *WeatherData
	GetAllWeatherDataReq chan bool
	GetAllWeatherDataRes chan map[string]*WeatherData
	WeatherDataList      map[string]*WeatherData `json:"weather_status_list"` // mysql checker list
}

var (
	Server *CacheServer
)

func NewCacheServer() *CacheServer {
	Server = &CacheServer{
		//http
		SetHttpData:       make(chan *HttpData),
		GetHttpDataReq:    make(chan string),
		GetHttpDataRes:    make(chan *HttpData),
		GetAllHttpDataReq: make(chan bool),
		GetAllHttpDataRes: make(chan map[string]*HttpData),
		HttpDataList:      map[string]*HttpData{},

		// mysql
		SetMysqlData:       make(chan *MysqlData),
		GetMysqlDataReq:    make(chan string),
		GetMysqlDataRes:    make(chan *MysqlData),
		GetAllMysqlDataReq: make(chan bool),
		GetAllMysqlDataRes: make(chan map[string]*MysqlData),
		MysqlDataList:      map[string]*MysqlData{},

		// weather
		SetWeatherData:       make(chan *WeatherData),
		GetWeatherDataReq:    make(chan string),
		GetWeatherDataRes:    make(chan *WeatherData),
		GetAllWeatherDataReq: make(chan bool),
		GetAllWeatherDataRes: make(chan map[string]*WeatherData),
		WeatherDataList:      map[string]*WeatherData{},
	}
	return Server
}

func (s *CacheServer) Run() {
	go func() {
		for {
			select {
			case data := <-s.SetWeatherData:
				s.WeatherDataList[data.LocationName] = data
			case locationName := <-s.GetWeatherDataReq:
				if data, exist := s.WeatherDataList[locationName]; exist {
					s.GetWeatherDataRes <- data
					continue
				}
				s.GetWeatherDataRes <- nil
			case <-s.GetAllWeatherDataReq:
				s.GetAllWeatherDataRes <- s.WeatherDataList
			}
		}
	}()
}
