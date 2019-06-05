package logwise

import (
	"bytes"
	"github.com/logwise/loglevel"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
)

var (
	StdLogger     = log.New(os.Stdout, "", 0)
	LogwiseLogger = log.New(os.Stdout, "", 0)
)

func TestLogger(t *testing.T) {
	l := Default()
	l.SetPrefix("wow")
	var buf bytes.Buffer

	l.SetOutput(&buf)
	l.Println("some text")
	if !strings.Contains(buf.String(), "wow") {
		t.Errorf("Prefix doesn't in it!")
	}

	buf.Reset()
	l.Trace("some text")
	if !strings.Contains(buf.String(), "TRACE") {
		t.Error("Not a \"TRACE\" log!")
	}

	buf.Reset()
	l.Debug("some text")
	if !strings.Contains(buf.String(), "DEBUG") {
		t.Error("Not a \"DEBUG\" log!")
	}

	buf.Reset()
	l.Info("some text")
	if !strings.Contains(buf.String(), "INFO") {
		t.Error("Not a \"INFO\" log!")
	}

	buf.Reset()
	l.Warn("some text")
	if !strings.Contains(buf.String(), "WARN") {
		t.Error("Not a \"WARNING\" log!")
	}

	buf.Reset()
	l.Error("some text")
	if !strings.Contains(buf.String(), "ERROR") {
		t.Error("Not a \"ERROR\" log!")
	}

	buf.Reset()
	l.Fatal("some text")
	if !strings.Contains(buf.String(), "FATAL") {
		t.Error("Not a \"FATAL\" log!")
	}

	buf.Reset()
	l.System("some text")
	if !strings.Contains(buf.String(), "SYS") {
		t.Error("Not a \"SYSTEM\" log!")
	}

	buf.Reset()
	l.SetLogLevel(loglevel.Error)
	l.Info("some text")
	if strings.Contains(buf.String(), "some text") {
		t.Error("Logging beyond log level!")
	}
}

func TestConcurrentSafe(t *testing.T) {
	l := Default()
	var buf bytes.Buffer
	l.SetOutput(&buf)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		for i := 0; i <= 100; i++ {
			l.Debugln("a")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i <= 100; i++ {
			l.Infoln("b")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i <= 100; i++ {
			l.Errorln("c")
		}
		wg.Done()
	}()

	wg.Wait()

	for _, line := range strings.Split(buf.String(), "\n") {
		if strings.Contains(line, "a") {
			if !strings.Contains(line, "DEBUG") {
				t.Error("Not concurrent safe!")
			}
		} else if strings.Contains(line, "b") {
			if !strings.Contains(line, "INFO") {
				t.Error("Not concurrent safe!")
			}
		} else if strings.Contains(line, "c") {
			if !strings.Contains(line, "ERROR") {
				t.Error("Not concurrent safe!")
			}
		}
	}
}

func BenchmarkStdLogger(b *testing.B) {
	StdLogger.Println("test")
}

func BenchmarkLogwiseLogger(b *testing.B) {
	LogwiseLogger.Println("test")
}
