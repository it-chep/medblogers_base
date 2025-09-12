package logger

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Graylog2/go-gelf/gelf"
	"github.com/joho/godotenv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// todo сделать вложенность в логах

type loggerKey string

const contextLoggerKey loggerKey = "clk"

// Logger - кастомная обертка над zap.Logger
type Logger struct {
	zap *zap.Logger
}

// todo доделать sidecar fluentbit

// New ...
func New() *Logger {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logHost := os.Getenv("LOG_HOST")
	logPort := os.Getenv("LOG_PORT")

	gelfWriter, err := gelf.NewWriter(fmt.Sprintf("%s:%s", logHost, logPort))
	if err != nil {
		panic(fmt.Sprintf("Ошибка создания логера GRAYLOG: %v", err))
	}

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "short_message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(gelfWriter),
		),
		zapcore.DebugLevel,
	)

	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	).With(
		zap.String("app", "medblogers_base"),
	)

	return &Logger{zap: logger}
}

// ContextWithLogger прокинуть логер в контекст
func ContextWithLogger(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, contextLoggerKey, l)
}

// fromContext взять логгер из контекста
func fromContext(ctx context.Context) *Logger {
	l, ok := ctx.Value(contextLoggerKey).(*Logger)
	if ok {
		return l
	}
	return nil
}

// Message записать в лог
func Message(ctx context.Context, format string) {
	l := fromContext(ctx)
	if l != nil {
		l.zap.Info(format)
	}
}

// Error записать ошибку
func Error(ctx context.Context, format string, err error) {
	l := fromContext(ctx)
	if l != nil {
		l.zap.Error(format, zap.Error(err))
	}
}
