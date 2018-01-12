package jsonlog

import (
	"fmt"
	"strings"
	"encoding/json"
	"errors"
)

type logmsg struct {
	logtype string
	timestamp string
	message string
	data interface{}
}

var isUnixTime bool = false
var hasSetTimeMode bool = false

// SetTimeMode sets whenever the time output should be UNIX or RFC3339 time.
func SetTimeMode(mode bool) {
	//Set mode, and set flag to ensure that logs are sucessful.
	isUnixTime = mode
	hasSetTimeMode = true

	//Notify other applications about time signature change.
	if mode {
		LogData(Info, "jsonlog_time_change", "UNIX")
		return
	}
	LogData(Info, "jsonlog_time_change", "RFC3339")
}

// Log prints out a log as json format, newlines will be converted to escaped \n.
func Log(t LogType, message string) error {
	//We just call LogData with no data.
	return LogData(t, message, nil)
}

// LogData prints out a log in json format, newlines will be converted to escaped \n.
// LogData allows attaching a data object which will be in a "data" tag.
func LogData(t LogType, message string, data interface{}) error {

	//Make sure that the mode is set, AND that software that might be listening
	//will know which mode to expect for time.
	if !hasSetTimeMode {
		return errors.New("Please set time mode before logging!")
	}
	
	//Escape \n
	tmsg := strings.Replace(message, "\n", "\\n", -1)
	
	//Turn type into string
	lt := GetTypeName(t)

	//Get current time string
	ti, err := TimeNow(isUnixTime).String()
	if err != nil {
		return err
	}

	//Create logmsg.
	l := logmsg{
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