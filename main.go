package main

import (
	"capi/app"
	"capi/logger"
	"fmt"
	"os"
)


func main() {
	fmt.Println(os.Getenv("TEST")) //env variable
	logger.Info("Starting app..........")
	app.Start()
}


