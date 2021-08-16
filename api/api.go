package api

import (
	"github.com/gin-gonic/gin"
)

func Api() {
	r := gin.Default()
	Router(r)
	r.Run(":8000") // listen and serve
}
