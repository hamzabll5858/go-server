package logger

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var Logger *zap.Logger = getLogger()

func getLogger() *zap.Logger {

	jsonFile, err := os.Open("config/log.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	fmt.Println("Successfully Opened log.json")
	defer jsonFile.Close()

	var cfg zap.Config

	if err := json.Unmarshal(byteValue, &cfg); err != nil {
		panic(err)
	}

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := cfg.Build()

	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return logger
}