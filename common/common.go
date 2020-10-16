package common

import (
	"encoding/json"
	"fmt"
	"futaapibooking/model"
	"io/ioutil"
	"os"
)

// var folder_setting string = "/go/bin/" // production
var folder_setting string = "" // devs

func GetConfigElasticSearch() model.ConfigElasticSearch {
	// Open our jsonFile
	jsonFile, err := os.Open(folder_setting + "setting/appconfig.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened appconfig.json")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	//var redisconfig RedisConnection
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	//appconfig, err := json.Marshal(result)
	appconfig, err := json.Marshal(result["ElasticSearch"])
	data := model.ConfigElasticSearch{}
	json.Unmarshal([]byte(appconfig), &data)

	return data
}
