package dao

import (
	"fmt"
	"server/model/dto"
)

func Websql(ID string) dto.WeatherRes {
	results, err := mysqlConn.Query("Select name,text FROM page3 where ID = ?", ID)
	if err != nil {
		fmt.Println(err.Error())
		return dto.WeatherRes{}
	}
	var pagedb dto.WeatherRes
	for results.Next() {

		err = results.Scan(&pagedb.Name, &pagedb.Text)
		if err != nil {
			fmt.Println(err.Error())
			return dto.WeatherRes{}
		}
		defer results.Close()

	}
	return pagedb
}
