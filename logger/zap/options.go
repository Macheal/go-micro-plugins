package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"

	"github.com/micro/go-micro/v2/logger"
)

type Options struct {
	logger.Options
	writeSyncer zapcore.WriteSyncer
}

type callerSkipKey struct{}

func WithCallerSkip(i int) logger.Option {
	return logger.SetOption(callerSkipKey{}, i)
}

type configKey struct{}

// WithConfig pass zap.Config to logger
func WithConfig(c zap.Config) logger.Option {
	return logger.SetOption(configKey{}, c)
}

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncoderConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) logger.Option {
	return logger.SetOption(encoderConfigKey{}, c)
}

type namespaceKey struct{}

func WithNamespace(namespace string) logger.Option {
	return logger.SetOption(namespaceKey{}, namespace)
}

type outwriter struct{}

func WithOutput(i io.Writer) logger.Option {
	return logger.SetOption(outwriter{}, i)
}

type multiOutput struct{}

func WithMultiOutput(i bool) logger.Option {
	return logger.SetOption(multiOutput{}, i)
}
