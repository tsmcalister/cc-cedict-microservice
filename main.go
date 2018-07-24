package main

import (
	"fmt"

	"github.com/tsmcalister/cc-cedict-microservice/service"
)

var appName = "cc-cedict-microservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("3141")
}
