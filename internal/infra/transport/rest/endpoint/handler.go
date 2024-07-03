package endpoint

import "github.com/gin-gonic/gin"

func NewRestHandler(router *gin.Engine) *Handler {
	return &Handler{
		router: router,
	}
}

type Handler struct {
	router *gin.Engine
}

func (h *Handler) Register() {
	basePath := "/api/v1"

	v1 := h.router.Group(basePath)
	{
		v1.GET("/ping", h.pingEndpoint)
	}
}
