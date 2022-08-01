package main

import (
	"fmt"

	"GoLangApi/app"
	"GoLangApi/logger"
)

func main() {

	fmt.Println("Api Started Listening ...")
	// log.Println("Statrted")
	logger.Info("Aplication Started..")
	app.Start()

}
