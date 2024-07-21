package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
)

func (rcv *Handler) Me(c *gin.Context) {
	token := c.MustGet("token").(*jwt.Token)
	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if userID, err := strconv.ParseUint(userIDString, 10, 64); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	} else {
		m, err := rcv.userStore.Get(uint(userID))
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.SecureJSON(http.StatusOK, gin.H{"username": m.Username})
	}
}
