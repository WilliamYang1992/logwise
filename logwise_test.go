package logwise

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	l := Default()
	l.Prefix = "wow"
	var buf bytes.Buffer

	l.Out = &buf
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
	if !strings.Contains(buf.String(), "SYSTEM") {
		t.Error("Not a \"SYSTEM\" log!")
	}

	buf.Reset()
	SetLogLevel(ErrorLevel)
	l.Info("some text")
	if strings.Contains(buf.String(), "some text") {
		t.Error("Logging beyond log level!")
	}
}
