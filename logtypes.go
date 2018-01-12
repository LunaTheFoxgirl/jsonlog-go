package jsonlog

// LogType is the type of log to be displayed.
type LogType uint8

const (

	// Info is an informational message.
	Info LogType(iota)

	// Warning is an warning message.
	Warning

	// Error is an error message.
	Error

	// Fatal is an fatal error message.
	Fatal
)

// typenames is a list of human readable names for the logtypes.
// Used during logging, so that the output isn't just numbers with no context.
var typenames []string = []string { "info", "warning", "error", "fatal"}

// Gets the typename belonging to the logtype.
func GetTypeName(t LogType) string {
	return typenames[int(t)]
}