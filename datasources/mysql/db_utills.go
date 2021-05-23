package mysql

import (
	"bitPay/logger"
	"encoding/json"
	"io/ioutil"
	"os"
)

func ParseJson(fileName string) (map[string]interface{}, error){
	jsonFile, err := os.Open(fileName)
	if err != nil {
		logger.Error("error when trying to parse json file", err)
		return nil, err
	}

	logger.Info("Successfully Opened json file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result, nil
}

