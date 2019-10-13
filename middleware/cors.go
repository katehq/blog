package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors cross region
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Access-Control-Allow-Origin"}
	config.AllowOrigins = []string{"http://192.168.1.225:8081","http://127.0.0.1:8081", "http://139.180.194.174"}
	config.AllowCredentials = true
	return cors.New(config)
}