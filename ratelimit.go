package ratelimit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Config is gin-ratelimit config
type Config struct {
	// Duration is limit interval, second
	Duration int64
	// RateLimit is limit count
	RateLimit int64
	// LimitFunc will be called when ip should be blocked
	LimitFunc func(c *gin.Context, ip string)
}

var defaultConfig = Config{
	Duration:  60,
	RateLimit: 60,
}

// DefaultConfig can create a default config
func DefaultConfig() Config {
	cp := defaultConfig
	return cp
}

// Default return a middleware with default config
func Default() gin.HandlerFunc {
	return New(defaultConfig)
}

// New return a middleware with config
func New(config Config) gin.HandlerFunc {
	rateLimiter := NewLimiter(config.Duration, config.RateLimit)
	if config.LimitFunc == nil {
		config.LimitFunc = func(c *gin.Context, ip string) {
			errorMsg := fmt.Sprintf("rate limit, requests should less than %d every %d seconds. ", config.RateLimit, config.Duration)
			c.JSON(http.StatusForbidden, gin.H{
				"ip":      ip,
				"message": errorMsg,
			})
		}
	}
	return func(c *gin.Context) {
		ip := c.ClientIP()
		shouldLimit := rateLimiter.ShouldLimit(ip)
		if shouldLimit {
			config.LimitFunc(c, ip)
			c.Abort()
		}
		c.Next()
	}
}
