package logging

import (
	"io"

	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levels = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

func New(cfg *config.Config, logWriter io.Writer) *zap.Logger {
	logLevel := zap.NewAtomicLevel()
	encoderCfg := encoderConfig()

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(zapcore.AddSync(logWriter)),
		logLevel,
	), zap.AddStacktrace(zapcore.ErrorLevel))

	logger = logger.Named(cfg.AppName)

	setLevel(cfg.LogLevel, &logLevel)
	logger = logger.With(zap.String("version", cfg.Version), zap.String("environment", cfg.Environment))
	return logger
}

func setLevel(level string, atom *zap.AtomicLevel) {
	if l, exists := levels[level]; exists {
		atom.SetLevel(l)
	}
}

func encoderConfig() zapcore.EncoderConfig {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.MessageKey = "message"
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.NameKey = "appName"

	return encoderCfg
}
