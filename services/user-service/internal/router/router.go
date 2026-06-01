package router

import (
	"ecommerce/pkg/logger"
	"ecommerce/user-service/internal/middleware"
	"ecommerce/user-service/internal/reqhandler"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func NewAuthRouter(db *bun.DB, logger *logger.Logger, rg *gin.RouterGroup) {
	h := reqhandler.NewHandler(db, logger)
	timeout := middleware.TimeoutMiddleware(time.Duration(time.Second * 7))
	rg.POST("/register", h.Register, timeout)
	rg.POST("/login", h.Login, timeout)
	profile := rg.Group("/profile", middleware.GetID())
	{
		profile.GET("/me", h.GetProfile, timeout)
	}
	logout := rg.Group("/logout", middleware.GetToken(), middleware.GetID(), timeout)
	{
		logout.POST("/single", h.Logout)
		logout.POST("/all", h.LogoutAll)
	}
}
