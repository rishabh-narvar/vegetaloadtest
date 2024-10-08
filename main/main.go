package main

import (
	"flag"
	"fmt"
	"time"
	"vegetaloadtest/parser"

	"vegetaloadtest/config"
	"vegetaloadtest/utils"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	var metrics vegeta.Metrics

	//conf:= config.InitConfig("../shopify/orders-qa.yaml")
	confArg := flag.String("config", "../config.yaml", "path to cofig yml")
	flag.Parse()
	fmt.Println(*confArg)
	//panic(0)
	conf := config.InitConfig(*confArg)

	//refactor this into constants file
	URL := conf.GetString("url")
	HTTP_METHOD := conf.GetString("httpmethod")
	VEGETA_RATE := conf.GetInt("rate")
	DURATION := conf.GetInt("duration")
	HEADERS := conf.GetStringMapString("headers")
	DYNAMIC_HEADERS := conf.GetStringMap("dynamic-headers")

	dynamic_headers := utils.ConvertToMapStringMapStringString(DYNAMIC_HEADERS)

	JSON_FILE_PATH := conf.GetString("post-request-json-file-path")
	DYNAMIC_FIELDS := conf.GetStringMapString("post-request-json-dynamic-fields")
	RESULTS_FILE_PATH := conf.GetString("dump-attack-results-file-path")
	REQUESTS_FILE_PATH := conf.GetString("dump-request-file-path")

	http_headers := utils.GetHttpHeaders(HEADERS)
	//   test_pacer := vegeta.SinePacer{Period: 1 * time.Second, Mean: vegeta.ConstantPacer{Freq: VEGETA_RATE, Per: time.Second}, Amp: }
	test_rate := vegeta.Rate{Freq: VEGETA_RATE, Per: time.Second}
	test_duration := time.Duration(DURATION) * (time.Second)

	requestsFileWriter, _ := utils.OpenFileCreateIfNotFound(REQUESTS_FILE_PATH)
	jsonString := parser.GetJsonString(JSON_FILE_PATH)
	fmt.Println(http_headers)
	// pacer := vegeta.SinePacer{
	// 	Period: 120 * time.Second,
	// 	Mean: vegeta.ConstantPacer{
	// 		Freq: 50,
	// 		Per:  time.Second,
	// 	},
	// 	Amp: vegeta.ConstantPacer{
	// 		Freq: 30,
	// 		Per: time.Second,
	// 	},
	// 	StartAt: math.Pi/6,
	// }
	targeter := utils.GetTargeter(URL, HTTP_METHOD, http_headers, jsonString, DYNAMIC_FIELDS, requestsFileWriter, dynamic_headers)
	attacker := vegeta.NewAttacker()
    response,_ := utils.OpenFileCreateIfNotFound("./response.bin")
	for res := range attacker.Attack(targeter, test_rate, test_duration, "Bang!") {
        fmt.Println("--attacking")
        res.Headers.Write(response)
        response.WriteString(fmt.Sprintf("%d", res.Code))
        response.WriteString(string(res.Body))
        response.WriteString("\n")
		metrics.Add(res)
	}
	reporter := vegeta.NewJSONReporter(&metrics)
	metrics.Close()

	utils.ProcessReport(reporter, RESULTS_FILE_PATH)
}
