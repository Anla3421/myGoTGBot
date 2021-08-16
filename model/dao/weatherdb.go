package dao

import (
	"fmt"
	"server/model/dto"

	_ "github.com/go-sql-driver/mysql"
)

func Weathersql(ID string) dto.WeatherRes {
	results, err := mysqlConn.Query("Select text FROM weather where ID = ?", ID)
	if err != nil {
		fmt.Println(err.Error())
		return dto.WeatherRes{}
	}
	var weadb dto.WeatherRes
	for results.Next() {

		err = results.Scan(&weadb.Text)
		if err != nil {
			fmt.Println(err.Error())
			return dto.WeatherRes{}
		}
		defer results.Close()

	}
	return weadb
}
