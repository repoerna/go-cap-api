package main

import (
	"capi/app"
	"capi/logger"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println(os.Getenv("Tets"))
	log.Printf("\nStarting application...\n")
	logger.Info("Starting application...(Zap log)")
	
	app.Start()
}