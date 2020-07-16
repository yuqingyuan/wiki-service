package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
)

var (
	database *gorm.DB

	username = "root"
	password = "root"
	dbName   = "wiki_crawler"
)

func init() {
	database, _ = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbName))
}

func Close() {
	if database != nil {
		if err := database.Close(); err != nil {
			fmt.Println(err)
		}
	}
}

func FetchEvents(month int, day int, offset int, limit int) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("date = '%d-%d'", month, day)
	rows, err := database.Table("events").Offset(offset).Limit(limit).Where(query).Rows()
	if err != nil {
		return nil, err
	}
	events := make([]map[string]interface{}, 0)
	for rows.Next() {
		event := Event{}
		err = rows.Scan(&event.ID, &event.Class, &event.Year, &event.Date, &event.Detail, &event.Links, &event.ImgLinks)
		if err != nil {
			return nil, err
		}

		jsonMap := make(map[string]interface{})
		jsonMap["id"] = event.ID
		jsonMap["type"] = event.Class
		jsonMap["year"] = event.Year
		jsonMap["date"] = event.Date
		jsonMap["detail"] = event.Detail
		jsonMap["links"] = event.Links
		jsonMap["images"] = strings.Split(event.ImgLinks, ",")

		events = append(events, jsonMap)
	}
	return events, nil
}
