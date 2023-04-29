package glog

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"
)

func Test(t *testing.T) {

	var buf bytes.Buffer

	l := Glog(Info)

	log.SetOutput(&buf)

	timestamp := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"

	infoMessage := "Info message"
	warnMessage := "Warn message"
	errorMessage := "Error message"

	l.Debug("Debug message")

	l.Info(infoMessage)
	l.Warn(warnMessage)
	l.Error(errorMessage)

	infoLog := timestamp + " " + "[INFO]" + " " + infoMessage
	warnLog := timestamp + " " + "[WARN]" + " " + warnMessage
	errorLog := timestamp + " " + "[ERROR]" + " " + errorMessage

	expected := infoLog + "\n" + warnLog + "\n" + errorLog + "\n"

	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Log output does not contain expected messages. Got:\n%s\nExpected:\n%s", buf.String(), expected)
	}
}