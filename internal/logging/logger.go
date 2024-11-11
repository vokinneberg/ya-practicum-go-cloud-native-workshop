package logging

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levels = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

func New(appName, level, environment, version string, logWriter io.Writer) *zap.Logger {
	logLevel := zap.NewAtomicLevel()
	cfg := encoderConfig()

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.Lock(zapcore.AddSync(logWriter)),
		logLevel,
	), zap.AddStacktrace(zapcore.ErrorLevel))

	logger = logger.Named(appName)

	setLevel(level, &logLevel)
	logger = logger.With(zap.String("version", version), zap.String("environment", environment))
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
