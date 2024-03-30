package api

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tanan/wg-in-handy/entity"
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
func (h *Handler) ShowWGInterface(c *gin.Context) {
	inf := h.Operator.ShowWGInterface()
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

func (h *Handler) getUserName(email string) string {
	// splitLen := 2
	s := strings.Split(email, "@")
	// if len(s) != splitLen {
	// 	return "", errors.New("email is not valid")
	// }
	return s[0]
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		slog.Error("cant't bind json", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"msg": "can't bind json"})
		return
	}
	if err := h.Operator.CreateUser(&entity.User{
		Name:  h.getUserName(user.Email),
		Email: user.Email,
	}); err != nil {
		slog.Error("can't create user", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "can't create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}

// TODO: download user config
func (h *Handler) DownloadUserConfig(c *gin.Context) {
	c.String(http.StatusOK, "DownloadUserConfig")
}
