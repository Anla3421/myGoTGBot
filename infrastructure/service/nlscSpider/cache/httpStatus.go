package cache

import (
	"time"
)

type HttpData struct {
	// *config.HttpConfig
	*HttpStatus
}

type HttpStatus struct {
	IsHealth  bool       `json:"is_health"`
	CheckTime *time.Time `json:"last_check_time"`
	Duration  int        `json:"duration"`
}

func NewHttpStatus() *HttpStatus {
	now := time.Now()
	return &HttpStatus{
		IsHealth:  false,
		Duration:  0,
		CheckTime: &now,
	}
}
