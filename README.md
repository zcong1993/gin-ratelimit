# gin-ratelimit [![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/gin-ratelimit)](https://goreportcard.com/report/github.com/zcong1993/gin-ratelimit) [![CircleCI branch](https://img.shields.io/circleci/project/github/zcong1993/gin-ratelimit/master.svg)](https://circleci.com/gh/zcong1993/gin-ratelimit/tree/master) [![codecov](https://codecov.io/gh/zcong1993/gin-ratelimit/branch/master/graph/badge.svg)](https://codecov.io/gh/zcong1993/gin-ratelimit)

> ratelimit middleware for gin

**IMPORTANT:** DO NOT USE IT IN PRODUCTION!!! USE [https://github.com/ulule/limiter](https://github.com/ulule/limiter)

## Install

```sh
$ go get -v github.com/zcong1993/gin-ratelimit
```

## Usage

```go
func main() {
    router := gin.New()
    router.Use(ratelimit.Default())
    router.GET("/", func(c *gin.Context) {
        c.String(200, testResp)
    })
    router.Run(":8080")
}
```

```go
func main() {
    router := gin.New()
    router.Use(ratelimit.New(ratelimit.Config{Duration: 1, RateLimit: 1}))
    router.GET("/", func(c *gin.Context) {
        c.String(200, testResp)
    })
    router.Run(":8080")
}
```

## License

MIT &copy; zcong1993
