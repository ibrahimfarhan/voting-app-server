package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

type Logger struct {
	Time     string `json:"time"`
	Level    string `json:"level"`
	Location string `json:"location,omitempty"`
	Message  string `json:"message,omitempty"`
	Err      string `json:"error,omitempty"`
}

func Info(message string) {
	l := Logger{Message: message, Level: "Info", Time: time.Now().Format("2006-01-02 15:04:05")}
	data, _ := json.Marshal(l)
	fmt.Println(string(data))
}

func Error(location string, err error) {
	l := Logger{Location: location, Err: err.Error(), Level: "Error", Time: time.Now().Format("2006-01-02 15:04:05")}
	data, _ := json.Marshal(l)
	fmt.Println(string(data))
}
