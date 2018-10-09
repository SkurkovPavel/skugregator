package skuParser


import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
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

func SaveData(data map[string]string) error {
	db, err := sql.Open("sqlite3", "/home/skurkov/GoProject/igor/parser_v2/db/notifications.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO notifications (number, method, platform, object, stage, date, nmc)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)", data["number"], data["method"], data["platform"], data["object"],
		data["stage"], data["date"], data["nmc"])
	if err != nil {
		return err
	}

	return nil
}


