package main

//
import "github.com/gin-gonic/gin"

//
func main() {

	//
	r := gin.Default()

	//
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	//
	r.GET("/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	//
	r.GET("/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	//
	r.GET("/search", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	//
	r.Run() // listen and serve on 0.0.0.0:8080...
}
