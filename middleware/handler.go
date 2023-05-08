package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Alfeenn/online-learning/helper"
	"github.com/Alfeenn/online-learning/model"
	"github.com/gin-gonic/gin"
)

func NewMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("X-API-KEY") == "RAHASIA" {
			return
		}

		ctx.Next()
	}
}

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := strconv.AppendBool([]byte(model.Key), true)
		claim := helper.ClaimToken(ctx, key)
		if claim.Username == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"msg": "UNAUTHORIZED", "code": http.StatusUnauthorized})
		}
		log.Print(gin.H{"data": claim.Username})
	}
}
