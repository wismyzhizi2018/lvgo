package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infof("this is %s", "我和我的祖国")
}
