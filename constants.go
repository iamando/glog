package welog

const (
	Error LogLevel = iota
	Warn
	Info
	Debug
)

const (
	Basic    LogMode = "basic"
	ColorFul LogMode = "colorful"
)

const TimestampFormat string = "2006-01-02 15:04:05"
