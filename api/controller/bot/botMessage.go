package bot

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type (
	MessageInput struct {
		Message string `json:"message" binding:"required"`
	}
	MessageResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"message,omitempty"`
		Result  interface{} `json:"result,omitempty"`
	}
)

func Message(c *gin.Context) {
	Input := MessageInput{}
	response := MessageResponse{}
	err := c.ShouldBindBodyWith(&Input, binding.JSON)
	if err != nil {
		fmt.Println("ShouldBindJSON fault", err)
		response.Status = 500
		response.Result = err.Error()

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = 200
	response.Message = Input.Message
	c.JSON(http.StatusOK, response)

	msg := tgbotapi.NewMessage(OwnerID, response.Message)
	msg.Text = response.Message
	BotConn.Send(msg)
}
