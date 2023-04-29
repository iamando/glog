package glog

import "io"

type LogLevel int

type Logger struct {
	level  LogLevel
	writer io.Writer
}
