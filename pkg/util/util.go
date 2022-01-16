package util

import (
	"encoding/json"
	"fmt"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/models"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"time"
)

func CheckError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
		return true
	}
	return false
}

func ConvertShanghaiTimeZone(t time.Time) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func ReadSettingsFromFile(settingFilePath string) (config models.Config) {
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}

func CreateTableIfNotExist(db *gorm.DB, tableModels []interface{}) {
	for _, value := range tableModels {
		if !db.HasTable(value) {
			db.CreateTable(value)
			fmt.Println("Create table ", reflect.TypeOf(value), " successfully")
		}
	}
}
