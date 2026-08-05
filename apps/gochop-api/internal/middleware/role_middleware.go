package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VendorOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		role := c.GetString("role")
		fmt.Println("========== DEBUG ==========")
		fmt.Println("ROLE:", role)
		fmt.Println("===========================")

		if role != "vendor" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Vendor access only",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		role := c.GetString("role")

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Admin access only",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ConsumerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		role := c.GetString("role")

		if role != "consumer" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Consumer access only",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
