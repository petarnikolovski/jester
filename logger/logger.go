package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func SetupLogger() {
	Log = logrus.New()

	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		Log.Panic(err)
	}
	Log.SetLevel(logLevel)

	Log.SetOutput(os.Stdout)
}
