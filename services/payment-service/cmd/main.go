package main

import (
	"ecommerce/pkg/logger"
	"log"
)

func main() {
	logger, err := logger.NewLogWriter("payment-service")
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Close()
}
