package loglevel

const (
	NotSet  LogLevel = 0
	Trace   LogLevel = 10
	Debug   LogLevel = 20
	Info    LogLevel = 30
	Warning LogLevel = 40
	Error   LogLevel = 50
	Fatal   LogLevel = 60
	System  LogLevel = 100
)

type LogLevel int

func (l LogLevel) String() string {
	var name string
	switch l {
	case NotSet:
		name = "NOTSET"
	case Trace:
		name = "TRACE"
	case Debug:
		name = "DEBUG"
	case Info:
		name = "INFO"
	case Warning:
		name = "WARNING"
	case Error:
		name = "ERROR"
	}
	return name
}

// GetPrefix return prefix for this log level
func (l LogLevel) GetPrefix() string {
	switch l {
	case NotSet:
		return "[NOTSET] "
	case Trace:
		return "[TRACE]  "
	case Debug:
		return "[DEBUG]  "
	case Info:
		return "[INFO]   "
	case Warning:
		return "[WARN]   "
	case Error:
		return "[ERROR]  "
	case Fatal:
		return "[FATAL]  "
	case System:
		return "[SYS]    "
	default:
		return "[NOTSET] "
	}
}

// ConvertStringToLogLevel convert input string to log level
func ConvertStringToLogLevel(logLevel string) LogLevel {
	var level LogLevel
	switch logLevel {
	case "NOTSET":
		level = NotSet
	case "TRACE":
		level = Trace
	case "DEBUG":
		level = Debug
	case "INFO":
		level = Info
	case "WARNING":
		level = Warning
	case "ERROR":
		level = Error
	case "FATAL":
		level = Fatal
	}
	return level
}
