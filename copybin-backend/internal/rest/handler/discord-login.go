package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rcv *Handler) DiscordLogin(c *gin.Context) {
	url := rcv.cfg.Discord.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}
