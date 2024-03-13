package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func VerifyToken(ctx *gin.Context) {
	getHeader := ctx.GetHeader("Authorization")

	split := strings.Split(getHeader, "Bearer ")

	errInvalidToken := errors.New("invalid token")

	if len(split) != 2 {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": errInvalidToken.Error(),
		})
		return
	}
	getToken := split[1]
	validated, err := jwt.Parse(getToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidToken
		}

		return []byte(viper.GetString("JWT_SECRET")), nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": errInvalidToken.Error(),
		})
		return
	}

	if _, ok := validated.Claims.(jwt.MapClaims); !ok && !validated.Valid {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": errInvalidToken.Error(),
		})

		return
	}

	ctx.Set("user", validated.Claims.(jwt.MapClaims))

	ctx.Next()

	// ctx.Set("user", validated.Claims.(jwt.MapClaims))

	// return validated.Claims.(jwt.MapClaims), nil
}
