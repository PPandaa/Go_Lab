package guard

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func getEncoder() zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(encoderConfig)

}

func getWriteSyncer() zapcore.WriteSyncer {

	writeSyncer, _, _ := zap.Open([]string{"stderr"}...)

	return writeSyncer

}

func init() {

	core := zapcore.NewCore(getEncoder(), getWriteSyncer(), zapcore.DebugLevel)
	Logger = zap.New(core, zap.AddCaller())

}
