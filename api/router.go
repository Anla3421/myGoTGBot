package api

import (
	"server/app/bot"
	"server/app/webpage"
	"server/view"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.LoadHTMLGlob("view/forms.html")

	r.GET("/form", view.Forms)
	r.POST("/form", webpage.Forms)
	r.POST("/message", bot.Message)
}
