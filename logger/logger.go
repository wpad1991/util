package logger

import (
	"encoding/json"
	"os"
	"time"

	"go.uber.org/zap"
)

var LogPath string
var logger *zap.Logger
var cfg zap.Config
var logmap map[string]*zap.Logger

func init() {
	logmap = make(map[string]*zap.Logger)
}

func SyncLogger() {
	for _, val := range logmap {
		val.Sync()
	}
}

func makelogger(folder string) *zap.Logger {
	LogPath = "/tmp/decol/log/" + folder + "/"
	if _, err := os.Stat(LogPath); os.IsNotExist(err) {
		os.MkdirAll(LogPath, os.ModePerm)
	}

	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/tmp/decol/logs"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	cfg.OutputPaths[1] = LogPath + time.Now().Format("2006-01-02,15_23_42")

	var err error
	logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	if logger == nil {
		panic("logger initailize logger is nil")
	}

	return logger
}

func GetLogger(folder string) *zap.Logger {
	if value, ok := logmap[folder]; ok {
		return value
	}
	return nil
}

func LogWrite(folder string, log interface{}) {
	if _, ok := logmap[folder]; !ok {
		logmap[folder] = makelogger(folder)
	}

	switch log.(type) {
	case error:
		logmap[folder].Error(log.(error).Error(), zap.String("datatime", time.Now().Format("2006-01-02 15:04:05")))
	case string:
		logmap[folder].Info(log.(string), zap.String("datatime", time.Now().Format("2006-01-02 15:04:05")))
	default:
		panic("LogWrite invalid log")
	}
}
