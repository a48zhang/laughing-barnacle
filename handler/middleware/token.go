package middleware

import (
	"github.com/gin-gonic/gin"
	"main/handler"
	"main/service"
)

func Validator() gin.HandlerFunc {
	return func(r *gin.Context) {
		tokenStr := r.GetHeader("Authorization")
		cl, err := service.ParseToken(tokenStr)
		if err != nil {
			handler.SendBadRequest(r, err, nil)
			r.Abort()
			return
		}
		r.Set("id", cl.Id)
		r.Next()
	}
}
