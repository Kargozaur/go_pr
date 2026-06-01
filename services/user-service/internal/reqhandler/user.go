package reqhandler

import (
	"context"
	"ecommerce/pkg/logger"
	"ecommerce/user-service/internal/jw"
	"ecommerce/user-service/internal/phasher"
	"ecommerce/user-service/internal/schemas"
	"ecommerce/user-service/internal/svc"
	"ecommerce/user-service/internal/thasher"
	"ecommerce/user-service/internal/util/validator"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type Handler struct {
	service *svc.UserService
	logger  *logger.Logger
}

func NewHandler(db *bun.DB, logger *logger.Logger) *Handler {
	phasher := phasher.NewPasswordHasher(12)
	thasher := thasher.NewHasher()
	validator := validator.NewValidator()
	iss := jw.NewJWT()
	service := svc.NewUserService(db, validator, phasher, thasher, iss)
	return &Handler{service: service, logger: logger}
}

func (h *Handler) Register(c *gin.Context) {
	ctx, cancel := h.prepareContext(c)
	defer cancel()
	var userBody schemas.RegisterSchema
	if err := c.ShouldBindJSON(&userBody); err != nil {
		h.logger.Writer.Error("bind fail", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Register(ctx, userBody); err != nil {
		h.logger.Writer.Error("register error", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "succesfully registred"})
}

func (h *Handler) Login(c *gin.Context) {
	ctx, cancel := h.prepareContext(c)
	defer cancel()
	var userBody schemas.LoginSchema
	if err := c.ShouldBindJSON(&userBody); err != nil {
		h.logger.Writer.Error("bind fail", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenPair, err := h.service.Login(ctx, userBody)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "request timeout"})
			return
		}
		h.logger.Writer.Error("login fail", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("access_token", tokenPair.AccessToken, 60*15, "/", os.Getenv("HOST"), true, true)
	c.SetCookie("refresh_token", tokenPair.RefreshToken, 60*60*24*7, "/", os.Getenv("HOST"), true, true)
	c.JSON(http.StatusOK, gin.H{"access_token": tokenPair.AccessToken, "token_type": tokenPair.TokenType})
}

func (h *Handler) LogoutAll(c *gin.Context) {
	ctx, cancel := h.prepareContext(c)
	defer cancel()
	userID, ok := c.Get("userID")
	if !ok {
		h.logger.Writer.Error("login fail", "error", "failed to get user id from middleware")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to read token"})
		return
	}
	if err := h.service.Logout(ctx, userID); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "request timeout"})
			return
		}
		h.logger.Writer.Error("failed to delete token from db", "error", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "succesfully logged out"})
}

func (h *Handler) Logout(c *gin.Context) {
	ctx, cancel := h.prepareContext(c)
	defer cancel()
	token, ok := c.Get("token")
	if !ok {
		h.logger.Writer.Error("login fail", "error", "failed to get user token from middleware")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to read token"})
		return
	}
	if err := h.service.Logout(ctx, token); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "request timeout"})
			return
		}
		h.logger.Writer.Error("failed to delete token from db", "error", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("access_token", "", 0, "/", os.Getenv("HOST"), true, true)
	c.SetCookie("refresh_token", "", 0, "/", os.Getenv("HOST"), true, true)

	c.JSON(http.StatusOK, gin.H{"success": "succesfully logged out"})
}

func (h *Handler) GetProfile(c *gin.Context) {
	ctx, cancel := h.prepareContext(c)
	defer cancel()
	userID, ok := c.Get("userID")
	if !ok {
		h.logger.Writer.Error("failed to get user id from middleware", "error", "user id not found")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}
	profile, err := h.service.GetProfile(ctx, userID)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "request timeout"})
			return
		}
		h.logger.Writer.Error("failed to get profile", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *Handler) prepareContext(c *gin.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(time.Second*7))
	return ctx, cancel
}
