package main

import (
	"capi/app"
	"capi/logger"
	"log"
)

func main() {
	log.Printf("\nStarting application...\n")
	logger.Info("Starting application...(Zap log)")
	
	app.Start()
}