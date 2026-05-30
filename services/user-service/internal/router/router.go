package router

import (
	"ecommerce/pkg/logger"
	"ecommerce/user-service/internal/middleware"
	"ecommerce/user-service/internal/reqhandler"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func NewAuthRouter(db *bun.DB, logger *logger.Logger, rg *gin.RouterGroup) {
	h := reqhandler.NewHandler(db, logger)
	rg.POST("/register", h.Register)
	rg.POST("/login", h.Login)
	logout := rg.Group("/logout", middleware.GetToken(), middleware.GetID())
	{
		logout.POST("/single", h.Logout)
		logout.POST("/all", h.LogoutAll)
	}
}
