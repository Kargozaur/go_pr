package main

import (
	"ecommerce/pkg/logger"
	"ecommerce/user-service/cfg"
	"ecommerce/user-service/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	logger, err := logger.NewLogWriter("user-service")
	if err != nil {
		log.Fatal(err)
	}
	db := cfg.NewDbConn()
	defer db.Close()
	defer logger.Close()
	app := gin.Default()
	r1 := app.Group("/users")
	{
		router.NewAuthRouter(db, logger, r1)
	}
	app.Run(":8082")
}
