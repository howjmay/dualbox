package main

import (
	"dualbox/cli"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
}

func main() {
	err := cli.App().Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
