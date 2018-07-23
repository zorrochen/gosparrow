package handler

import (
	"github.com/gin-gonic/gin"
	. "gosparrow/base/http_server"
	"net/http"
)

func Init() {
	Register("GET", "/ping", ping)
}

// ping
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
