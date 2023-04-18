package logger

import (
	"log"

	"go.uber.org/zap"
)

func InitialLogger() {
	conf := zap.NewDevelopmentConfig()
	conf.DisableStacktrace = true
	logger, err := conf.Build()
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(logger)
}
