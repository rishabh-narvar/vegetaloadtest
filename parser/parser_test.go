package parser

import "testing"

func TestGetJsonString(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetJsonString(tt.args.filePath); got != tt.want {
				t.Errorf("GetJsonString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPreparedJsonForRequest(t *testing.T) {
	type args struct {
		jsonString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPreparedJsonForRequest(tt.args.jsonString, map[string]string{}); got != tt.want {
				t.Errorf("GetPreparedJsonForRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
