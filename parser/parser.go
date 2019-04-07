package parser

import (
	"github.com/tidwall/sjson"
	"io/ioutil"
    "os"
    "github.com/google/uuid"
)


func GetJsonString(filePath string) string{
    jsonFile, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
    jsonString := string(byteValue[:])
    
    return jsonString
}

func GetPreparedJsonForRequest(jsonString string, dynamicfields []string) string  {
    newJSON:= jsonString   

    for _, field := range dynamicfields {
		newJSON, _= sjson.Set(jsonString, field , uuid.New().String())
	}

    return newJSON
}