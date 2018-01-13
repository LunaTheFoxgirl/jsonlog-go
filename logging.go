package jsonlog

import (
	"fmt"
	"strings"
	"encoding/json"
)

type logmsg struct {
	UNIX string `json:"unix"`
	LogType string `json:"type"`
	Timestamp string `json:"timestamp"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// Log prints out a log as json format, newlines will be converted to escaped \n.
func Log(t LogType, message string) error {
	//We just call LogData with no data.
	return LogData(t, message, nil)
}

// LogData prints out a log in json format, newlines will be converted to escaped \n.
// LogData allows attaching a data object which will be in a "data" tag.
func LogData(t LogType, message string, data interface{}) error {
	//Escape \n
	tmsg := strings.Replace(message, "\n", "\\n", -1)
	
	//Turn type into string
	lt := GetTypeName(t)

	//Get current time string
	ti, tu, err := TimeNow().String()
	if err != nil {
		return err
	}

	//Create logmsg.
	l := logmsg{
		tu,
		lt,
		ti,
		tmsg,
		data,
	}
	
	//Marshal it into json.
	b, err := json.Marshal(l)
	if err != nil {
		return err
	}
	
	//Print to screen.
	fmt.Println(string(b))
	return nil
}