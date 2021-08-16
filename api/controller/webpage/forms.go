package webpage

import (
	"fmt"

	"server/api/controller/bot"
	"server/view"

	"github.com/gin-gonic/gin"
)

func Forms(c *gin.Context) {
	input := bot.Input{}
	if err := c.ShouldBind(&input); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
		return
	}
	fmt.Printf("Input Message:\nEmail: %s\nSubject: %s\nMessage: %s\n", input.Email, input.Subject, input.Message)
	bot.LoginInfo(input)
	view.Forms(c)
}
