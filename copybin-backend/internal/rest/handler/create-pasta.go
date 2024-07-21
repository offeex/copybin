package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type createPastaRequest struct {
	text string
}

func (rcv *Handler) CreatePasta(c *gin.Context) {
	var req createPastaRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	m, err := rcv.pastaStore.Create(req.text, 1)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to create pasta"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pasta created", "id": m.ID})
}
