/*
 * @Author: lihuan
 * @Date: 2021-09-30 14:38:52
 * @LastEditTime: 2021-10-02 08:57:34
 * @Email: 17719495105@163.com
 */
package middleware

import (
	"go-lh/constants"
	"go-lh/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var code int

		code = constants.SUCCESS

		Authorization := ctx.GetHeader("Authorization")

		token := strings.Split(Authorization, " ")

		// 需要 token
		if Authorization == "" {
			code = constants.INVALID_TOKEN_ERROR
		} else {
			_, err := utils.ParseToken(token[1])
			if err != nil {

				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = constants.TOKEN_TIMEOUT_ERROR
				default:
					code = constants.TOKEN_FAIL
				}
			}
		}

		if code != constants.SUCCESS {
			utils.Response(ctx, http.StatusUnauthorized, code, nil)

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
