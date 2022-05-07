package main

//
import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

//
func main() {

	//
	r := gin.Default()

	//
	r.Use(favicon.New("./favicon.ico")) // set favicon middleware

	//
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello... Welcome... Presentation...",
		})
	})

	//
	r.GET("/help", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Help... Manual...",
		})
	})

	//
	r.GET("/all_emojis", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List all emojis...",
		})
	})

	//
	r.GET("/all_categories", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List all categories' name and descprition...",
		})
	})

	//
	r.GET("/all_sub_categories", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List all categories' name, descprition and category to which it belongs...",
		})
	})

	//
	r.Run() // listen and serve on 0.0.0.0:8080...
}
