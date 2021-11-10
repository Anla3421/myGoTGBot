package cache

import (
	"server/service/nlscSpider/config"
	"time"
)

type MysqlData struct {
	*config.MysqlConfig
	*MysqlStatus
}

type MysqlStatus struct {
	ConnNum   int        `json:"conn_num"`
	IsHealth  bool       `json:"is_health"`
	CheckTime *time.Time `json:"last_check_time"`
	Duration  int        `json:"duration"`
}

func NewMysqlStatus() *MysqlStatus {
	now := time.Now()
	return &MysqlStatus{
		ConnNum:   0,
		IsHealth:  false,
		Duration:  0,
		CheckTime: &now,
	}
}
