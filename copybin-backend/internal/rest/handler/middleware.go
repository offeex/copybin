package handler

import (
	"copybin/pkg"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (rcv *Handler) AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		_ = c.AbortWithError(http.StatusUnauthorized, errors.New("empty auth header"))
		return
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 {
		_ = c.AbortWithError(
			http.StatusBadRequest,
			errors.New("invalid auth header format"),
		)
		return
	}

	if authHeaderParts[0] != "Bearer" {
		_ = c.AbortWithError(
			http.StatusBadRequest,
			errors.New("unsupported auth type"),
		)
		return
	}

	tokenString := authHeaderParts[1]
	token, err := pkg.VerifyToken(tokenString, rcv.cfg.JWTSecret)
	if err != nil || !token.Valid {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.Set("token", token)
	c.Next()
}
