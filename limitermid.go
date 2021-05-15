package limiter

import (
	"github.com/gin-gonic/gin"
	"log"
)


func LimiterWithMaxconn(maxConn int) gin.HandlerFunc{
	limiter:= newLimiter(maxConn)
	return limitermiddle(limiter)
}

func limitermiddle(l *limiter) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		err:=l.GetCon()
		if err != nil {
			log.Println("The visitor too much!")
			ctx.Abort()
		}
		ctx.Next()
		l.Release()
	}
}
