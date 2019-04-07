package config

import (
	"reflect"
	"testing"
	"github.com/spf13/viper"
)

func TestInitConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *viper.Viper
	}{
		
		{
			name: "with path",
			args: args{path: "../config.yaml"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitConfig(tt.args.path); reflect.TypeOf(got) != reflect.TypeOf(viper.GetViper())  {
				t.Errorf("InitConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseConfig(t *testing.T) {
	type args struct {
		byteArray []byte
	}
	tests := []struct {
		name string
		args args
		expected string
    expectedheader string
	}{
		{
			name: "With byte Array",
			args: args{byteArray: []byte(`
url: google.com
httpmethod: POST
headers:
  - 'content-type: text/plain'
body: 
  path: ./request_body.json

`)},
			expected: "google.com",
      expectedheader: "content-type: text/plain",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
      got := parseConfig(tt.args.byteArray);
			if  got.Get("url") != tt.expected {
				t.Errorf("parseConfig() = %v, want %v", got, tt.expected)
			}

      headers:= got.GetStringSlice("headers")
      for _, v := range headers {
        //fmt.Println(v)
        if(v != tt.expectedheader){
          t.Errorf("headers() = %v, want", got)
        }
      }
		})
	}
}
