package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginWeb := gin.Default()

	ginWeb.GET("/test", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "golang",
			"tag":  "<br>",
		}

		context.AsciiJSON(http.StatusOK, data)
	})
	ginWeb.Run(":8080")
}
