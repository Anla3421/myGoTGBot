package dao

import (
	"fmt"
	"server/domain/repository/dto"
)

func Drinksql(Drinkid int, Who string, Arguments string, Sugar string, Ice string) {
	results, err := mysqlConn.Query("INSERT INTO drink (ID,who,drink,sugar,ice) VALUES(?,?,?,?,?) ON DUPLICATE KEY UPDATE who=?,drink=?,sugar=?,ice=?",
		Drinkid, Who, Arguments, Sugar, Ice, Who, Arguments, Sugar, Ice)
	if err != nil {
		panic(err)

	}
	defer results.Close()
}

func Drinksqlget(ID int) dto.DrinkRes {
	results, err := mysqlConn.Query("SELECT * FROM drink where ID=?", ID)
	if err != nil {
		fmt.Println(err.Error())
		return dto.DrinkRes{}
	}
	var drinkdb dto.DrinkRes
	for results.Next() {

		err = results.Scan(&drinkdb.ID, &drinkdb.Who, &drinkdb.Drink, &drinkdb.Sugar, &drinkdb.Ice)
		if err != nil {
			fmt.Println(err.Error())
			return dto.DrinkRes{}
		}
		defer results.Close()

	}
	return drinkdb
}

func Drinksqltruncate() {
	results, err := mysqlConn.Query("TRUNCATE table drink")
	if err != nil {
		panic(err.Error())

	}
	defer results.Close()
}
