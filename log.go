package log

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	js := `{"level": "DEBUG", "encoding": "json", "outputPaths": ["stdout"], "errorOutputPaths": ["stdout"]}`
	cfg := &zap.Config{}
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	cfg.DisableCaller = true
	cfg.DisableStacktrace = false
	logger, _ = cfg.Build()

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
}

func Info(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Infow(msg, keysAndValues...)

}

func Debug(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Debugw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Errorw(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Warnw(msg, keysAndValues...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Panicw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Fatalw(msg, keysAndValues...)
}

func String(key string, value string) zap.Field {
	return zap.String(key, value)
}

func FieldErr(err error) zap.Field {
	return zap.Error(err)
}

func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}
