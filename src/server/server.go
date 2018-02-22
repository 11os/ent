package server

import (
	"../utils"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type TV struct {
	id int32
	url string
}

const tvs = `[{"id":1,"url":"https://www.dygod.net/html/tv/oumeitv/109955.html"},{"id":2,"url":"https://www.dygod.net/html/tv/oumeitv/109675.html"}]`

func Start() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"shield":  utils.Run("https://www.dygod.net/html/tv/oumeitv/109955.html"),
			"sheldon": utils.Run("https://www.dygod.net/html/tv/oumeitv/109675.html"),
		})
	})
	r.GET("/tv/:id", func(c *gin.Context) {
		id := c.Param("id")
		url := gjson.Get(tvs, `#[id="`+ id +`"].url`).String()
		// utils.Print("url ============== " + url)
		c.JSON(200, utils.Run(url))
		// c.JSON(200, url)
	})
	r.GET("/tvs", func(c *gin.Context) {
		c.JSON(200, tvs)
	})
	r.Run(":13000") // listen and serve on 0.0.0.0:8080
}
