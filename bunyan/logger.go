package bunyan

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	pid = os.Getpid()
)

const (
	LOG_FATAL   = 100
	LOG_EMERG   = 90
	LOG_ALERT   = 80
	LOG_CRIT    = 70
	LOG_ERR     = 60
	LOG_WARNING = 50
	LOG_NOTICE  = 40
	LOG_INFO    = 30
	LOG_DEBUG   = 20
)

type Logger struct {
	Name string

	logger *log.Logger
}

func NewLogger(name string) Logger {
	logger := log.New(os.Stdout, "", 0)
	return Logger{Name: name, logger: logger}
}

func (logger *Logger) getBaseData(data map[string]interface{}) map[string]interface{} {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["name"] = logger.Name
	data["pid"] = pid
	hostname, err := os.Hostname()
	if err == nil {
		data["hostname"] = hostname
	}
	data["time"] = time.Now()
	return data
}

func (l *Logger) Log(level int, v ...interface{}) error {
	data := l.getBaseData(nil)
	data["msg"] = fmt.Sprint(v...)
	data["level"] = level
	out, _ := json.Marshal(data)
	l.logger.Println(string(out))

	return nil
}

func (l *Logger) Logm(level int, s string, m map[string]interface{}) error {
	data := l.getBaseData(m)
	data["msg"] = s
	data["level"] = level
	out, _ := json.Marshal(data)
	l.logger.Println(string(out))

	return nil
}

// func (l *Logger) InfoData(d map[string]interface{}) error {
// l.logger.Println("Info")
// return nil
// }

// NORMAL LEVELS
func (l *Logger) Debug(v ...interface{}) error {
	return l.Log(LOG_DEBUG, v...)
}

func (l *Logger) Info(v ...interface{}) error {
	return l.Log(LOG_INFO, v...)
}

func (l *Logger) Warning(v ...interface{}) error {
	return l.Log(LOG_WARNING, v...)
}

func (l *Logger) Err(v ...interface{}) error {
	return l.Log(LOG_ERR, v...)
}

// FATAL
func (l *Logger) Fatal(v ...interface{}) (err error) {
	err = l.Log(LOG_FATAL, v...)
	os.Exit(1)
	return
}

func (l *Logger) Debugf(s string, v ...interface{}) error {
	return l.Log(LOG_DEBUG, fmt.Sprintf(s, v...))
}

func (l *Logger) Infof(s string, v ...interface{}) error {
	return l.Log(LOG_INFO, fmt.Sprintf(s, v...))
}

func (l *Logger) Warningf(s string, v ...interface{}) error {
	return l.Log(LOG_WARNING, fmt.Sprintf(s, v...))
}

func (l *Logger) Errf(s string, v ...interface{}) error {
	return l.Log(LOG_ERR, fmt.Sprintf(s, v...))
}

func (l *Logger) Fatalf(s string, v ...interface{}) error {
	l.Log(LOG_FATAL, fmt.Sprintf(s, v...))
	os.Exit(1)
	return nil
}

func (l *Logger) Debugm(s string, m map[string]interface{}) error {
	return l.Logm(LOG_DEBUG, s, m)
}

func (l *Logger) Infom(s string, m map[string]interface{}) error {
	return l.Logm(LOG_INFO, s, m)
}

func (l *Logger) Warningm(s string, m map[string]interface{}) error {
	return l.Logm(LOG_WARNING, s, m)
}

func (l *Logger) Errm(s string, m map[string]interface{}) error {
	return l.Logm(LOG_ERR, s, m)
}

func (l *Logger) Fatalm(s string, m map[string]interface{}) error {
	l.Logm(LOG_FATAL, s, m)
	os.Exit(1)
	return nil
}
