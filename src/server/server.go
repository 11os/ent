package server

import (
	"../utils"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"shield":  utils.Run("https://www.dygod.net/html/tv/oumeitv/109955.html"),
			"sheldon": utils.Run("https://www.dygod.net/html/tv/oumeitv/109675.html"),
		})
	})
	r.Run(":13000") // listen and serve on 0.0.0.0:8080
}
