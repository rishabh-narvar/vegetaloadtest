package main

import (
	"strings"
	"net/http"
	"github.com/perf/parser"
	"fmt"
	"os"
  "time"
  "perf/config"
  vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
  conf:= config.InitConfig("../config.yaml")

  URL := conf.GetString("url")
  HTTPMETHOD := conf.GetString("httpmethod")
  VEGETARATE := conf.GetInt("rate")
  DURATION := conf.GetInt("duration")
  HEADERS := conf.GetStringMapString("headers")
  JSONFILEPATH := conf.GetString("json-body-file-path")
  DYNAMICFIELDS := conf.GetStringSlice("json-body-dynamic-fields")
  
  jsonString := parser.GetJsonString(JSONFILEPATH)
  

  //fmt.Println(jsonStringForRequest)

  rate := vegeta.Rate{Freq: VEGETARATE, Per: time.Second}
  duration := time.Duration(DURATION) * (time.Second)

  tr := getTargeter(URL, HTTPMETHOD, HEADERS, jsonString, DYNAMICFIELDS)
  attacker := vegeta.NewAttacker()

  var metrics vegeta.Metrics
  for res := range attacker.Attack(tr, rate, duration, "Bang!") {
    metrics.Add(res)
  }

  reporter := vegeta.NewTextReporter(&metrics)
  metrics.Close()
  fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
  fmt.Printf("report %s", reporter.Report(os.Stdout))
}

func getTargeter (url string, httpmethod string, headers map[string]string, body string, dynamicFields []string) vegeta.Targeter{
  return func() vegeta.Targeter {
    return func(t *vegeta.Target) (err error) {
        t.Method = httpmethod
        t.URL    = url
        if body != "" && strings.ToUpper(httpmethod) != "GET" {
          jsonStringForRequest := parser.GetPreparedJsonForRequest(body, dynamicFields)
          fmt.Println(jsonStringForRequest)
          t.Body = []byte(jsonStringForRequest)
        }
        t.Header = http.Header{}
        
        for headerKey, headerValue := range headers{
          t.Header.Add(headerKey, headerValue)
        }
        return err
    }
  }()
}

