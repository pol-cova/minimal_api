package mapi

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func SetupLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		logrus.SetOutput(logFile)
	} else {
		logrus.Warn("Failed to log to file, using default stderr")
	}
}

func Logger(c *Context) {
	startTime := time.Now()
	c.Next()
	duration := time.Since(startTime)
	logrus.WithFields(logrus.Fields{
		"method":      c.Method(),
		"path":        c.Path(),
		"status_code": c.Response.StatusCode(),
		"duration":    duration,
	}).Info("Request handled")
}
