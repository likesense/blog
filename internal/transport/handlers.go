package http

import (
	"blog/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	*AuthorizationHandler
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		AuthorizationHandler: NewAuthorizationHandlers(services),
	}
}

func (h *Handler) RegisterAPI(router *gin.RouterGroup) {
	api := router.Group("/api")
	h.RegisterUserAPI(api)
}

func (h *Handler) RegisterUserAPI(grp *gin.RouterGroup) {
	user := grp.Group("/user")
	{
		h.AuthorizationHandler.RegisterNewUserHandler(user)
	}
}
