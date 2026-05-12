package main

import (
	"ecommerce/pkg/logger"
	"log"
)

func main() {
	logger, err := logger.NewLogWriter("gateway")
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Close()
}
