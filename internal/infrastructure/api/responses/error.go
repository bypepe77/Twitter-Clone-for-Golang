package responses

import "github.com/gin-gonic/gin"

func Error(statusCode int, c *gin.Context, message string) {
	c.JSON(statusCode, gin.H{
		"message": message,
	})
}
