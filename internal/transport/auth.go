package http

import (
	"blog/internal/models"
	"blog/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorizationHandler struct {
	Authorization services.Authorization
	JWT           services.JWT
}

func NewAuthorizationHandlers(services *services.Services) *AuthorizationHandler {
	return &AuthorizationHandler{
		Authorization: services.Authorization,
		JWT:           services.JWT,
	}
}

type signInInput struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ah *AuthorizationHandler) RegisterNewUserHandler(grp *gin.RouterGroup) {
	user := grp.Group("")
	{
		user.POST("/sign-up", ah.SingUp)
		user.POST("/sign-in", ah.SignIn)
	}
}

func (ah *AuthorizationHandler) SignIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't bind user data"})
		return
	}

	user, err := ah.Authorization.GetUserCred(input.Nickname, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't get user credentials"})
		return
	}
	accessToken, err := ah.JWT.GenerateAccessToken(user.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't generate access token"})
	}

	refreshToken, err := ah.JWT.GenerateAccessToken(user.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't generate refresh token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (ah *AuthorizationHandler) SingUp(ctx *gin.Context) {
	var input models.User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	id, err := ah.Authorization.CreateUserCred(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't create user"})
		return
	}
	accessToken, err := ah.JWT.GenerateAccessToken(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't generate access token"})
		return
	}
	refreshToken, err := ah.JWT.GenerateRefreshToken(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't generate refresh token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":            id,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
