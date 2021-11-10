package selfTime

import (
	"errors"
	"time"

	"server/service/mylib/errorCode"
)

const (
	TimeLayout = "2006-01-02 15:04:05"
	DateLayout = "2006-01-02"
)

var (
	// 台灣時區
	TWLoc *time.Location
)

func init() {
	TWLoc, _ = time.LoadLocation("Asia/Taipei")
}

type HourMinuteSecond struct {
	Hour   int // 時
	Minute int // 分
	Second int // 秒
}

// 取得一天的起始時間
func StartOfDate(t *time.Time) *time.Time {
	start_time := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, TWLoc)
	return &start_time
}

// 取得一天的起始時間(指定的timezone)
func StartOfDateWithTimeZone(t *time.Time, loc *time.Location) *time.Time {
	start_time := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
	return &start_time
}

// 取得一天的結束時間
func EndOfDate(t *time.Time) *time.Time {
	end_time := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 59, TWLoc)
	return &end_time
}

// 取得一天的結束時間(指定的timezone)
func EndOfDateWithTimeZone(t *time.Time, loc *time.Location) *time.Time {
	end_time := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 59, loc)
	return &end_time
}

// 轉換秒數成時分秒格式
func TransSecondToHourMinuteSecond(sec int) (data *HourMinuteSecond) {
	duration := time.Duration(sec) * (time.Second)
	data = &HourMinuteSecond{
		Hour:   int(duration.Hours()),
		Minute: (int(duration.Minutes()) % 60),
		Second: int((duration % time.Minute).Seconds()),
	}
	return
}

// 轉換時分秒成秒數格式
func TransHourMinuteSecondToSecond(time *HourMinuteSecond) (sec int) {
	sec = time.Second
	if time.Hour != 0 {
		sec += time.Hour * 60 * 60
	}
	if time.Minute != 0 {
		sec += time.Minute * 60
	}
	return
}

// 確認時間是否早於現在
func CheckIsTimeEarlierThanNow(t *time.Time) (code int, data interface{}, err error) {
	now := time.Now()
	if !t.Before(now) {
		return errorCode.TimeCantBeLaterThanNow, nil, errors.New("time is later than now")
	}
	return
}

// 確認時間是否晚於現在
func CheckIsTimeLaterThanNow(t *time.Time) (code int, data interface{}, err error) {
	now := time.Now()
	if !t.After(now) {
		return errorCode.TimeCantBeEarlierThanNow, nil, errors.New("time is earlier than now")
	}
	return
}

// 確認開始時間及結束時間的格式(duration 不為nil時 代表會檢查endTime - startTime的區間不能大於duration)
func CheckStartTimeAndEndTime(startTime *time.Time, endTime *time.Time, duration *time.Duration) (code int, data interface{}, err error) {
	if startTime == nil {
		return errorCode.TimeSettingError, nil, errors.New("Start time is nil")
	}

	if endTime == nil {
		return errorCode.TimeSettingError, nil, errors.New("End time is nil")
	}

	if startTime.After(*endTime) {
		return errorCode.TimeSettingError, nil, errors.New("end time can't be later than start time")
	}

	if duration != nil {
		if endTime.Sub(*startTime) > *duration {
			return errorCode.TimeRangeIsTooLog, nil, errors.New("duration between start and end time is too long")
		}
	}

	return
}
