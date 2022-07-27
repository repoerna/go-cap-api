package main

import (
	"capi/app"
	"capi/logger"
)

func main() {
	logger.Info("starting the application ...")
	app.Start()
}
