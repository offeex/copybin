package handler

import (
	"copybin/pkg"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rcv *Handler) DiscordCallback(c *gin.Context) {
	// Verify state
	state := c.Query("state")
	if state != "state" {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("invalid state"))
		return
	}

	// Exchange code for access and refresh tokens
	code := c.Query("code")
	codeRes, err := rcv.cfg.Discord.Exchange(c, code)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Retrieve user data
	authClient := rcv.cfg.Discord.Client(c, codeRes)
	userDataRes, err := authClient.Get("https://discord.com/api/users/@me")
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else if userDataRes.StatusCode != http.StatusOK {
		_ = c.AbortWithError(http.StatusUnauthorized, errors.New("invalid codeRes"))
		return
	}

	defer func() {
		if userDataRes.Body.Close() != nil {
			return
		}
	}()

	// Parse user data
	var userInfo struct {
		Id       json.Number `json:"id"`
		Username string      `json:"username"`
	}
	if err := json.NewDecoder(userDataRes.Body).Decode(&userInfo); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Create user if doesn't exist
	userModel, err := rcv.userStore.CreateIfNotExists(userInfo.Username)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Create integration if it doesn't exist
	if err := rcv.integrationStore.CreateIfNotExists(
		userModel.ID,
		"discord",
		userInfo.Id.String(),
		codeRes.AccessToken,
		codeRes.RefreshToken,
	); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Construct JWT
	jwtToken, err := pkg.CreateToken(userModel.ID, rcv.cfg.JWTSecret)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Authorization", "Bearer "+jwtToken)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
