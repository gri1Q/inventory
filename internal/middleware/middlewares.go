package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), duration)
		defer cancel()

		// Заменяем контекст в запросе
		c.Request = c.Request.WithContext(ctx)

		// Канал для сигнала о завершении запроса
		done := make(chan struct{})

		// Запускаем обработку запроса в отдельной горутине
		go func() {
			defer close(done)
			c.Next()
		}()

		select {
		case <-done:
			// Запрос завершился нормально
			return
		case <-ctx.Done():
			// Таймаут истёк — прерываем запрос
			c.JSON(http.StatusGatewayTimeout, gin.H{
				"error":   "request timeout",
				"message": "запрос выполнялся слишком долго",
			})
			c.Abort()
			return
		}
	}
}

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

		c.Next()
	}
}
func MyRoutCustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			// Останавливаем запрос и возвращаем 405
			c.JSON(405, gin.H{"error": "Ошибка на маршруте " + c.Request.RequestURI})
			c.Abort() // прекращает дальнейшую обработку
			return
		}

		c.Next()
	}
}
