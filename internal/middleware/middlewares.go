package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MyGlobalCustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			// Останавливаем запрос и возвращаем 405
			c.JSON(405, gin.H{"error": "Глобальный мидлвар "})
			c.Abort() // прекращает дальнейшую обработку
			return
		}

		c.Next() // только GET запросы проходят дальше
	}
}

func MyGroupCustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			// Останавливаем запрос и возвращаем 405
			c.JSON(405, gin.H{"error": "Групповой мидлвар"})
			c.Abort() // прекращает дальнейшую обработку
			return
		}

		c.Next() // только GET запросы проходят дальше
	}
}
func MyRoutCustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			// Останавливаем запрос и возвращаем 405
			c.JSON(405, gin.H{"error": "Ошибка на маршруте " + c.Request.RequestURI})
			c.Abort() // прекращает дальнейшую обработку
			return
		}

		c.Next() // только GET запросы проходят дальше
	}
}
