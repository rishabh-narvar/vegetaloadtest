package utils

import (
	"errors"
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

func GetTargeter (url string, httpmethod string, headers http.Header, body string, dynamicFields map[string]string, requestDumpFileWriter *os.File) vegeta.Targeter{
    return func() vegeta.Targeter {
      return func(t *vegeta.Target) (err error) {
          t.Method = httpmethod
          t.URL    = url
          if body != "" && strings.ToUpper(httpmethod) != "GET" {
            jsonStringForRequest := parser.GetPreparedJsonForRequest(body, dynamicFields)
            t.Body = []byte(jsonStringForRequest)

            //hack determine a better way to do this. Have to dig into vegeta docs, if there is a handle to requests object in targetter
            //spin off a go routine now
            if requestDumpFileWriter != nil {
                go requestDumpFileWriter.WriteString(jsonStringForRequest + " \n")
            }
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

  func OpenFileCreateIfNotFound(filePath string)(*os.File, error){
    if filePath == "" {
        fmt.Errorf("Invalid file path %s ", filePath)
        return nil, errors.New("Invalid file path")
    }
   return os.OpenFile(filePath, os.O_RDWR | os.O_CREATE, 0777)
  }

  func ProcessReport(reporter vegeta.Reporter, filePath string){
    //debug
    fmt.Printf("report %s", reporter.Report(os.Stdout))
    file, err := OpenFileCreateIfNotFound(filePath)

    if err != nil{
        fmt.Errorf("Error dumping into results file %s ", err)
        return
      
      DumpReportToFile(reporter, file)
    }
  }