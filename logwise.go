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
	"github.com/WilliamYang1992/logwise/loglevel"
	"io"
	"log"
	"os"
	"sync"
)

// 默认 flag
const DefaultFlags = log.Ldate | log.Ltime | log.Lshortfile

// 默认输出
var DefaultOutput = os.Stdout

type Logger struct {
	prefix string
	out    io.Writer
	mu     sync.Mutex
	level  loglevel.LogLevel
	logger *log.Logger
}

func (l Logger) String() string {
	return fmt.Sprintf("Logger<Prefix: %s, Out: %s>", l.prefix, l.out)
}

func (l Logger) Prefix() string {
	return l.prefix
}

func (l Logger) FullPrefix(lv loglevel.LogLevel) string {
	var ret string
	if l.prefix != "" {
		ret = l.prefix + " " + lv.GetPrefix()
	} else {
		ret = lv.GetPrefix()
	}
	return ret
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

func (l Logger) LogLevel() loglevel.LogLevel {
	return l.level
}

func (l *Logger) SetLogLevel(lv loglevel.LogLevel) {
	l.level = lv
}

// 按照不同的 loglevel 改变 logger 属性
func (l *Logger) dress(lv loglevel.LogLevel) {
	l.logger.SetPrefix(l.FullPrefix(lv))
	if l.out == DefaultOutput {
		if lv != loglevel.System && lv >= loglevel.Error {
			l.logger.SetOutput(os.Stderr)
		} else {
			l.logger.SetOutput(l.out)
		}
	} else {
		l.logger.SetOutput(l.out)
	}
}

func (l *Logger) print(lv loglevel.LogLevel, v ...interface{}) {
	if lv == loglevel.System || lv >= l.level {
		l.mu.Lock()
		l.dress(lv)
		_ = l.logger.Output(3, fmt.Sprint(v...))
		l.mu.Unlock()
	}
}

func (l *Logger) println(lv loglevel.LogLevel, v ...interface{}) {
	if lv == loglevel.System || lv >= l.level {
		l.mu.Lock()
		l.dress(lv)
		_ = l.logger.Output(3, fmt.Sprintln(v...))
		l.mu.Unlock()
	}

}

func (l *Logger) printf(lv loglevel.LogLevel, format string, v ...interface{}) {
	if lv == loglevel.System || lv >= l.level {
		l.mu.Lock()
		l.dress(lv)
		_ = l.logger.Output(3, fmt.Sprintf(format, v...))
		l.mu.Unlock()
	}
}

// Default, based on the log level of this logger
func (l *Logger) Print(v ...interface{}) {
	l.print(l.level, v...)
}

func (l *Logger) Println(v ...interface{}) {
	l.println(l.level, v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.printf(l.level, format, v...)
}

// NOTSET
func (l *Logger) NotSet(v ...interface{}) {
	l.print(loglevel.NotSet, v...)
}

func (l *Logger) NotSetln(v ...interface{}) {
	l.println(loglevel.NotSet, v...)
}

func (l *Logger) NotSetf(format string, v ...interface{}) {
	l.printf(loglevel.NotSet, format, v...)
}

// TRACE
func (l *Logger) Trace(v ...interface{}) {
	l.print(loglevel.Trace, v...)
}

func (l *Logger) Traceln(v ...interface{}) {
	l.println(loglevel.Trace, v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.printf(loglevel.Trace, format, v...)
}

// DEBUG
func (l *Logger) Debug(v ...interface{}) {
	l.print(loglevel.Debug, v...)
}

func (l *Logger) Debugln(v ...interface{}) {
	l.println(loglevel.Debug, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.printf(loglevel.Debug, format, v...)
}

// INFO
func (l *Logger) Info(v ...interface{}) {
	l.print(loglevel.Info, v...)
}

func (l *Logger) Infoln(v ...interface{}) {
	l.println(loglevel.Info, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.printf(loglevel.Info, format, v...)
}

// WARNING
func (l *Logger) Warn(v ...interface{}) {
	l.print(loglevel.Warning, v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	l.println(loglevel.Warning, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.printf(loglevel.Warning, format, v...)
}

// ERROR
func (l *Logger) Error(v ...interface{}) {
	l.print(loglevel.Error, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.println(loglevel.Error, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.printf(loglevel.Error, format, v...)
}

// FATAL
func (l *Logger) Fatal(v ...interface{}) {
	l.print(loglevel.Fatal, v...)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.println(loglevel.Fatal, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.printf(loglevel.Fatal, format, v...)
}

// SYSTEM
func (l *Logger) System(v ...interface{}) {
	l.print(loglevel.System, v...)
}

func (l *Logger) Systemln(v ...interface{}) {
	l.println(loglevel.System, v...)
}

func (l *Logger) Systemf(format string, v ...interface{}) {
	l.printf(loglevel.System, format, v...)
}

// New logger
func New(out io.Writer, prefix string, flag int) *Logger {
	logger := new(Logger)
	logger.out = out
	logger.prefix = prefix
	logger.logger = log.New(out, "", flag)
	return logger
}

// Default logger
func Default() *Logger {
	return New(DefaultOutput, "", DefaultFlags)
}
