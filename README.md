# GinLimiter

A limiter middleware for gin web framework <br>
一个适用于Gin框架的流控中间件


# Install
```
go get -u github.com/l0n3rs/GinLimiter
```

# Useage
```
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/l0n3rs/GinLimiter"
)

func main(){
	server:=gin.Default()
	maxConn:=100
	server.Use(limiter.LimiterWithMaxconn(maxConn))
	server.GET("/", func(ctx *gin.Context) {
		ctx.String(200,"Limiter using")
	})
	server.Run(":8080")
}

```
