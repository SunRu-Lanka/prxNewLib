package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

//init adding loger function
//using Ubar Zap logger
func init() {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey=""
	config.EncoderConfig = encoderConfig
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}
// Info crate  logger function
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

//Debug logger function
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

//Error logger function
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
