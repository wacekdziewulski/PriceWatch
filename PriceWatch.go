package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	logrus.SetLevel(logrus.DebugLevel)

	application := Initialize()

	application.start()
}
