package main

import (
	"github.com/vad/go-bunyan/bunyan"
)

var (
	logger = bunyan.NewLogger("example")
)

func main() {
	logger.Info("Test message ", "by me")
	logger.Infof("Test message n %d", 1)
}
