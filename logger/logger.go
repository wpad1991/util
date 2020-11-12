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

func init() {
	LogPath = "/tmp/decol/log/"
	if _, err := os.Stat(LogPath); os.IsNotExist(err) {
		os.MkdirAll(LogPath, os.ModePerm)
	}

	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/tmp/atm/logs"],
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

}

func GetLogger() *zap.Logger {
	return logger
}

func LogWrite(log interface{}) {

	switch log.(type) {
	case error:
		logger.Error(log.(error).Error(), zap.Time("datatime", time.Now()))
	case string:
		logger.Info(log.(string), zap.Time("datatime", time.Now()))
	default:
		panic("LogWrite invalid log")
	}
}
