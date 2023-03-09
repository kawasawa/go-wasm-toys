package logging

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Log(msgID messageID, err error, options ...zap.Field)
}

type Logger struct{}

var logger ILogger = newLogger()

func newLogger() *Logger {
	return &Logger{}
}

func GetLogger() ILogger {
	return logger
}

func (l *Logger) Log(msgID messageID, err error, options ...zap.Field) {
	msg, ok := messages[msgID]
	if !ok {
		msg = message{zapcore.ErrorLevel, "MessageID not found."}
	}
	if err != nil {
		fmt.Println("["+strings.ToUpper(msg.level.String())+"]", msg.message, err.Error())
	} else {
		fmt.Println("["+strings.ToUpper(msg.level.String())+"]", msg.message)
	}
}
