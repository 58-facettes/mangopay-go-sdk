package log

import "log"

// Logger is an interface for using your own Logging tool like Zap or Logrus.
type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
}

var DefaultLogger = &Log{}

func NewLogger() *Log {
	return &Log{}
}

// Log is just a dummy log tool that you can use.
// for more advenced usage please use Zap or Logrus.
type Log struct {
}

func (l *Log) Debug(args ...interface{}) {
	log.Print(args...)
}
func (l *Log) Debugf(template string, args ...interface{}) {
	log.Printf(template, args...)
}
func (l *Log) Error(args ...interface{}) {
	log.Print(args...)
}
func (l *Log) Errorf(template string, args ...interface{}) {
	log.Printf(template, args...)
}
func (l *Log) Fatal(args ...interface{}) {
	log.Print(args...)
}
func (l *Log) Fatalf(template string, args ...interface{}) {
	log.Printf(template, args...)
}
func (l *Log) Warnf(template string, args ...interface{}) {
	log.Printf(template, args...)
}
