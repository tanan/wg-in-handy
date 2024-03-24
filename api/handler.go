package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanan/wg-in-handy/operator"
)

type Handler struct {
	Operator *operator.Operator
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// TODO: show wg configuration
func (h *Handler) ShowInterface(c *gin.Context) {
	inf := h.Operator.ShowInterface()
	c.JSON(http.StatusOK, inf)
}

// TODO: get user list
func (h *Handler) GetUsers(c *gin.Context) {
	c.String(http.StatusOK, "GetUsers")
}

// TODO: get user info by userid
func (h *Handler) GetUser(c *gin.Context) {
	c.String(http.StatusOK, "GetUser")
}

func (h *Handler) CreateUser(c *gin.Context) {
	c.String(http.StatusOK, "CreateUser")
}

// TODO: download user config
func (h *Handler) DownloadUserConfig(c *gin.Context) {
	c.String(http.StatusOK, "DownloadUserConfig")
}
