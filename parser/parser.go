package parser

import (
	"time"
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

func GetPreparedJsonForRequest(jsonString string, dynamicfields map[string]string) string  {
    newJSON:= jsonString   

    for key, value := range dynamicfields {
        switch value{
            case "timestamp":
               newJSON, _ = sjson.Set(newJSON, key, time.Now().UTC().Format(time.RFC3339))
               //fmt.Println(newJSON)
            case "uuid":
                newJSON, _= sjson.Set(newJSON, key , uuid.New().String())
            case "epoch":
                newJSON, _= sjson.Set(newJSON, key , time.Now().Unix())
            case "epochnano":
                newJSON, _= sjson.Set(newJSON, key , time.Now().UnixNano())
        }
    }
    return newJSON
}