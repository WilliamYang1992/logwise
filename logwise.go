/*
MIT License

Copyright (c) 2019 William Yang

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package logwise

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	TraceLevel   LogLevel = 0
	DebugLevel   LogLevel = 10
	InfoLevel    LogLevel = 20
	WarningLevel LogLevel = 30
	ErrorLevel   LogLevel = 40
	FatalLevel   LogLevel = 50
	SystemLevel  LogLevel = 100
)

// 默认 flag
const DefaultFlags = log.Ldate | log.Ltime | log.Lshortfile

// 日志过滤等级
var level LogLevel

// 默认输出
var DefaultOutput = os.Stdout

type LogLevel int

func (l LogLevel) String() string {
	var name string
	switch l {
	case TraceLevel:
		name = "TRACE"
	case DebugLevel:
		name = "DEBUG"
	case InfoLevel:
		name = "INFO"
	case WarningLevel:
		name = "WARNING"
	case ErrorLevel:
		name = "ERROR"
	}
	return name
}

type Logger struct {
	prefix string
	out    io.Writer
	mu     sync.Mutex
	logger *log.Logger
}

func (l Logger) String() string {
	return fmt.Sprintf("Logger<Prefix: %s, Out: %s>", l.prefix, l.out)
}

func (l Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l Logger) Output() io.Writer {
	return l.out
}

func (l *Logger) SetOutput(out io.Writer) {
	l.out = out
}

// 按照不同的 loglevel 改变 logger 属性
func (l *Logger) dress(lv LogLevel) {
	l.logger.SetPrefix(GetPrefix(lv, l.prefix))
	if l.out == DefaultOutput {
		if lv != SystemLevel && lv >= ErrorLevel {
			l.logger.SetOutput(os.Stderr)
		}
	} else {
		l.logger.SetOutput(l.out)
	}
}

func (l *Logger) print(lv LogLevel, v ...interface{}) {
	if lv == SystemLevel || lv >= level {
		l.mu.Lock()
		l.dress(lv)
		_ = l.logger.Output(3, fmt.Sprint(v...))
		l.mu.Unlock()
	}
}

func (l *Logger) println(lv LogLevel, v ...interface{}) {
	if lv == SystemLevel || lv >= level {
		l.mu.Lock()
		l.dress(lv)
		_ = l.logger.Output(3, fmt.Sprintln(v...))
		l.mu.Unlock()
	}

}

func (l *Logger) printf(lv LogLevel, format string, v ...interface{}) {
	if lv == SystemLevel || lv >= level {
		l.mu.Lock()
		l.dress(lv)
		_ = l.logger.Output(3, fmt.Sprintf(format, v...))
		l.mu.Unlock()
	}
}

// Default, based on global log level
func (l *Logger) Print(v ...interface{}) {
	l.print(level, v...)
}

func (l *Logger) Println(v ...interface{}) {
	l.println(level, v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.printf(level, format, v...)
}

// TRACE
func (l *Logger) Trace(v ...interface{}) {
	l.print(TraceLevel, v...)
}

func (l *Logger) Traceln(v ...interface{}) {
	l.println(TraceLevel, v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.printf(TraceLevel, format, v...)
}

// DEBUG
func (l *Logger) Debug(v ...interface{}) {
	l.print(DebugLevel, v...)
}

func (l *Logger) Debugln(v ...interface{}) {
	l.println(DebugLevel, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.printf(DebugLevel, format, v...)
}

// INFO
func (l *Logger) Info(v ...interface{}) {
	l.print(InfoLevel, v...)
}

func (l *Logger) Infoln(v ...interface{}) {
	l.println(InfoLevel, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.printf(InfoLevel, format, v...)
}

// WARNING
func (l *Logger) Warn(v ...interface{}) {
	l.print(WarningLevel, v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	l.println(WarningLevel, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.printf(WarningLevel, format, v...)
}

// ERROR
func (l *Logger) Error(v ...interface{}) {
	l.print(ErrorLevel, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.println(ErrorLevel, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.printf(ErrorLevel, format, v...)
}

// FATAL
func (l *Logger) Fatal(v ...interface{}) {
	l.print(FatalLevel, v...)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.println(FatalLevel, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.printf(FatalLevel, format, v...)
}

// SYSTEM
func (l *Logger) System(v ...interface{}) {
	l.print(SystemLevel, v...)
}

func (l *Logger) Systemln(v ...interface{}) {
	l.println(SystemLevel, v...)
}

func (l *Logger) Systemf(format string, v ...interface{}) {
	l.printf(SystemLevel, format, v...)
}

// New logger
func New(out io.Writer, prefix string, flag int) *Logger {
	logger := new(Logger)
	logger.out = out
	logger.prefix = prefix
	logger.logger = log.New(out, "", flag)
	return logger
}

// 默认 logger
func Default() *Logger {
	return New(DefaultOutput, "", DefaultFlags)
}

// 获取当前的 loglevel
func GetLogLevel() LogLevel {
	return level
}

// 设置当前的 loglevel
func SetLogLevel(l LogLevel) {
	level = l
}

// 转换成 LogLevel
func ConvertStringToLogLevel(logLevel string) LogLevel {
	var level LogLevel
	switch logLevel {
	case "TRACE":
		level = TraceLevel
	case "DEBUG":
		level = DebugLevel
	case "INFO":
		level = InfoLevel
	case "WARNING":
		level = WarningLevel
	case "ERROR":
		level = ErrorLevel
	case "FATAL":
		level = FatalLevel
	}
	return level
}

// 获取 Logger 前缀
func GetPrefix(level LogLevel, prefix string) string {
	if prefix != "" {
		prefix += " "
	}
	switch level {
	case TraceLevel:
		prefix += "[TRACE]   "
	case DebugLevel:
		prefix += "[DEBUG]   "
	case InfoLevel:
		prefix += "[INFO]    "
	case WarningLevel:
		prefix += "[WARN]    "
	case ErrorLevel:
		prefix += "[ERROR]   "
	case FatalLevel:
		prefix += "[FATAL]   "
	case SystemLevel:
		prefix += "[SYSTEM]  "
	}
	return prefix
}
