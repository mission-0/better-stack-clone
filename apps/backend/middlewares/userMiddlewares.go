package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Usermiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Token")

		token := strings.Split(header, " ")
		fmt.Println("token is ", token[1])
		// 		fmt.Println("Header token is:", header)
		fmt.Println("Usermiddleware Func")

		jwtToken, err := jwt.Parse(token[1], func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			fmt.Println("error while decoding ", err)
		}
		if jwtToken.Valid {

			claimsMaps := (jwtToken.Claims).(jwt.MapClaims) // it will return a map of claims

			fmt.Println("token parsed signature", claimsMaps["Id"])
			fmt.Println("token is valid")

			ctx.Set("userId", claimsMaps["Id"])
			userId, err := ctx.Get("userId")
			if !err {
				fmt.Println("Error getting userId from gin context")
			}

			fmt.Println("userId from Gin context is : ", userId)
			ctx.Next()
		}

		if !jwtToken.Valid {
			fmt.Println("token is invalid")
			return
		}
	}
}
