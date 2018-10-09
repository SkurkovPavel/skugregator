package skuParser

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var SkuDB *sql.DB
//Все что связано с базой

//Получаем результаты

//Формируем запросы

//Записываем в базу

//Сообщаем о результатах


func GetData(notification string) (map[string]string, error) {
	db, err := sql.Open("sqlite3", "/home/skurkov/GoProject/igor/parser_v2/db/notifications.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var number, method, platform, object, stage, date, nmc string
	data := map[string]string{}

	row := db.QueryRow("SELECT * FROM notifications WHERE number = $1", notification)
	err = row.Scan(&number, &method, &platform, &object, &stage, &date, &nmc)
	if err != nil{
		return nil, err
	}

	data["number"] = number
	data["method"] = method
	data["platform"] = platform
	data["object"] = object
	data["stage"] = stage
	data["date"] = date
	data["nmc"] = nmc

	return data, nil
}

func SaveData(data map[string]string,) error {


	var err error
		_,err = SkuDB.Exec("UPDATE responses SET body = $1, time=time() WHERE site = $2", data["body"],data["site"])

	if err != nil {
		return err
	}

	return nil
}

func OnenSkuDB() error {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbPath := string(pwd)+"/skugregator/alias/skuDataBase.db"

	SkuDB, err = sql.Open("sqlite3", dbPath)

	if err != nil {
		return err
	}
	return nil
}

func CloseSkuDB(db *sql.DB){
	db.Close()
}

