package main

import (
	"capi/app"
	"capi/logger"
	"fmt"
	"log"
	"os"
)

func sanityCheck(){
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
	}

	for _, envKey := range envProps {
		if os.Getenv(envKey) == "" {
			logger.Fatal(fmt.Sprintf("environtment variabel %s not defined. terminating application..", envKey))
		}
	}
}

func main() {
	fmt.Println(os.Getenv("Tets"))
	log.Printf("\nStarting application...\n")
	logger.Info("Starting application...(Zap log)")
	
	app.Start()
}