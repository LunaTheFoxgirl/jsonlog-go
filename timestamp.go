package jsonlog

import (
	"time"
	"strconv"
)

// Timestamp is a point in time.
type Timestamp struct {
	t time.Time
	UnixTime bool
}

// TimeNow creates a new timestamp for the current time.
func TimeNow(isUnix bool) Timestamp {
	return Timestamp{ time.Now(), isUnix }
}

// String returns the timestamp as a string
func (t Timestamp) String() (string, error) {
	if t.UnixTime {
		return strconv.FormatInt(t.t.Unix(), 10), nil
	}
	s, err := t.t.MarshalText()
	if err != nil {
		return "", err
	}
	return string(s), nil
}