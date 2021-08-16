package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Text struct {
	Email   string
	Subject string
	Message string
}

func Forms(c *gin.Context) {
	c.HTML(http.StatusOK, "forms.html", nil)
}
