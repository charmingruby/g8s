package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Pong bool `json:"pong"`
}

func (h *Handler) pingEndpoint(c *gin.Context) {
	res := PingResponse{
		Pong: true,
	}

	c.JSON(http.StatusOK, res)
}
