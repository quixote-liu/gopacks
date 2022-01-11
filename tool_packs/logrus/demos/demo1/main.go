package main

import "github.com/sirupsen/logrus"

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	logger.Debug("hello ", "debug")
	logger.Info("hello ", "hubei")
	logger.Warn("hello ", "warn")
	logger.Error("hello ", "error")
	logger.Fatal("hello ", "fatal")
	logger.Panic("hello ", "panic")
}
