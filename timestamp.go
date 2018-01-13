package jsonlog

import (
	"time"
	"strconv"
)

// Timestamp is a point in time.
type Timestamp struct {
	t time.Time
}

// TimeNow creates a new timestamp for the current time.
func TimeNow() Timestamp {
	return Timestamp{ time.Now() }
}

// String returns the timestamp as a string.
// first value is the RFC3339 version.
// second value is UNIX time version.
func (t Timestamp) String() (string, string, error) {
	s, err := t.t.MarshalText()
	if err != nil {
		return "", "", err
	}
	return string(s), strconv.FormatInt(t.t.Unix(), 10), nil
}