package loggerutil

import (
    "github.com/sirupsen/logrus"
)

type Logger struct {
    Log *logrus.Logger
}

func NewLogger() *Logger {
    log := logrus.New()

    log.SetLevel(logrus.InfoLevel)
    log.SetFormatter(&logrus.TextFormatter{
        ForceColors:     true,
        FullTimestamp:   true,
        TimestampFormat: "2006-01-02 00:00:00",
    })

    return &Logger{
        Log: log,
    }
}