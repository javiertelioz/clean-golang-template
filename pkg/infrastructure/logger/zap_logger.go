package logger

import (
	"fmt"
	"runtime"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/javiertelioz/clean_architecture/pkg/domain/contracts/services"
)

type logMessage struct {
	level   string
	message string
	caller  string
}

type ZapLogger struct {
	logger    *zap.Logger
	msgChan   chan logMessage
	wg        sync.WaitGroup
	closeOnce sync.Once
	done      chan struct{}
}

func NewLogger() services.LoggerService {
	config := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			MessageKey:    "message",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize zap logger: %v", err))
	}

	l := &ZapLogger{
		logger:  logger,
		msgChan: make(chan logMessage, 1000), // Buffered channel
		done:    make(chan struct{}),
	}

	l.startWorker()
	return l
}

func (z *ZapLogger) startWorker() {
	z.wg.Add(1)
	go func() {
		defer z.wg.Done()
		for {
			select {
			case msg := <-z.msgChan:
				switch msg.level {
				case "trace", "debug":
					z.logger.Debug(msg.message, zap.String("loc", msg.caller))
				case "info":
					z.logger.Info(msg.message, zap.String("loc", msg.caller))
				case "warn":
					z.logger.Warn(msg.message, zap.String("loc", msg.caller))
				case "error":
					z.logger.Error(msg.message, zap.String("loc", msg.caller), zap.Stack("stacktrace"))
				}
			case <-z.done:
				return
			}
		}
	}()
}

func (z *ZapLogger) Close() {
	z.closeOnce.Do(func() {
		close(z.done)
		z.wg.Wait()
		_ = z.logger.Sync()
	})
}

func (z *ZapLogger) getCallerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func (z *ZapLogger) log(level, msg string) {
	select {
	case z.msgChan <- logMessage{level: level, message: msg, caller: z.getCallerInfo()}:
	default:
		z.logger.Warn("logging channel full, logging synchronously",
			zap.String("level", level),
			zap.String("message", msg))
	}
}

func (z *ZapLogger) Trace(msg string) {
	z.log("trace", msg)
}

func (z *ZapLogger) Debug(msg string) {
	z.log("debug", msg)
}

func (z *ZapLogger) Info(msg string) {
	z.log("info", msg)
}

func (z *ZapLogger) Warn(msg string) {
	z.log("warn", msg)
}

func (z *ZapLogger) Error(msg string) {
	z.log("error", msg)
}
