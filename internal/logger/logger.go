package logger

import (
	"github.com/ilhame90/gymshark-task/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	logger *zap.SugaredLogger
}

var logLevels = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func New(conf *config.Config) (*zap.Logger, error) {
	zapConf := zap.NewProductionConfig()
	zapConf.Development = conf.Env == "development"

	if level, exists := logLevels[conf.Log.Level]; exists {
		zapConf.Level = zap.NewAtomicLevelAt(level)
	} else {
		zapConf.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
