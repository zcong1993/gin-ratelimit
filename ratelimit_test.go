package ratelimit

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func init() {
	gin.SetMode(gin.TestMode)
}

const testResp = "test resp"

func newServer(mw gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(mw)
	router.GET("/", func(c *gin.Context) {
		c.String(200, testResp)
	})
	return router
}

func makeReq(s *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	s.ServeHTTP(w, r)
	return w
}

func TestDefaultConfig(t *testing.T) {
	c := DefaultConfig()
	assert.Equal(t, c.Duration, int64(60))
	assert.Equal(t, c.RateLimit, int64(60))
}

func TestMw(t *testing.T) {
	mw := New(Config{Duration: 1, RateLimit: 1})
	s := newServer(mw)
	w := makeReq(s)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), testResp)
	w = makeReq(s)
	assert.Equal(t, w.Code, 403)
	assert.Equal(t, w.Body.String(), "{\"ip\":\"\",\"message\":\"rate limit, requests should less than 1 every 1 seconds. \"}\n")
	time.Sleep(time.Second)
	w = makeReq(s)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), testResp)
}

func TestCustomLimitFunc(t *testing.T) {
	testLimitResp := "rate limit"
	mw := New(Config{Duration: 1, RateLimit: 1, LimitFunc: func(c *gin.Context, ip string) {
		c.String(500, testLimitResp)
	}})
	s := newServer(mw)
	w := makeReq(s)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), testResp)
	w = makeReq(s)
	assert.Equal(t, w.Code, 500)
	assert.Equal(t, w.Body.String(), testLimitResp)
	time.Sleep(time.Second)
	w = makeReq(s)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), testResp)
}

func TestDefault(t *testing.T) {
	mw := Default()
	s := newServer(mw)
	for i := int64(0); i < defaultConfig.RateLimit; i++ {
		w := makeReq(s)
		assert.Equal(t, w.Code, 200)
		assert.Equal(t, w.Body.String(), testResp)
	}
	w := makeReq(s)
	assert.Equal(t, w.Code, 403)
}
