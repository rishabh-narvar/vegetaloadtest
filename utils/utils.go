package utils

import (
	"os"
	"fmt"
	"io"
	"strings"
    "net/http"
    vegeta "github.com/tsenart/vegeta/lib"
    "github.com/perf/parser"
)

func GetHttpHeaders(headers map[string]string) http.Header{
    header := http.Header{}
    for headerKey, headerValue := range headers{
        header.Add(headerKey, headerValue)
    }
    return header
}

func GetTargeter (url string, httpmethod string, headers http.Header, body string, dynamicFields map[string]string) vegeta.Targeter{
    return func() vegeta.Targeter {
      return func(t *vegeta.Target) (err error) {
          t.Method = httpmethod
          t.URL    = url
          if body != "" && strings.ToUpper(httpmethod) != "GET" {
            jsonStringForRequest := parser.GetPreparedJsonForRequest(body, dynamicFields)
            t.Body = []byte(jsonStringForRequest)
          }
          t.Header = headers
          return err
      }
    }()
  }

  func DumpReportToFile(reporter vegeta.Reporter, writer io.Writer){
    if writer != nil {
      err := reporter.Report(writer)
      if err != nil {
          fmt.Errorf("Error %s", err)
      }
    }
  }

  func ProcessReport(reporter vegeta.Reporter, filePath string){
    //debug
    fmt.Printf("report %s", reporter.Report(os.Stdout))
  
    if filePath != "" {
      file, err := os.OpenFile(filePath, os.O_RDWR | os.O_CREATE, 0777)
      if err != nil{
        fmt.Errorf("Error dumping into results file %s ", err)
        return
      }
      DumpReportToFile(reporter, file)
    }
  }