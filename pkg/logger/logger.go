package logger

import "github.com/sirupsen/logrus"

func NewLogger() *logrus.Logger {
  logger := logrus.New()
	// logger.SetFormatter(&logrus.JSONFormatter{
	// 	PrettyPrint: true,
	// })
	// logger.WithFields(logrus.Fields{
	// 	"key": "value",
	// }).Info("This is JSON Formatter with additional fields.")
  return logger
}
