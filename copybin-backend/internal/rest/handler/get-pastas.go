package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rcv *Handler) GetPastas(c *gin.Context) {
	pastas, err := rcv.pastaStore.ListAll(10)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to get pastas"},
		)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"pastas": pastas})
}
