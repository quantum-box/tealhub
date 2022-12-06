package log

import "go.uber.org/zap"

// Logger has zap logging interface
var Logger *zap.SugaredLogger

func init() {
	l, _ := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	defer func() {
		_ = l.Sync()
	}()

	Logger = l.Sugar()
}

// SetDebug switch debug mode for logger
func SetDebug() {
	l, _ := zap.NewDevelopment()
	defer func() {
		_ = l.Sync()
	}()

	Logger = l.Sugar()
}

