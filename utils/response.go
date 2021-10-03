/*
 * @Author: lihuan
 * @Date: 2021-09-24 08:27:57
 * @LastEditTime: 2021-09-30 17:56:10
 * @Email: 17719495105@163.com
 */
package utils

import (
	"go-lh/constants"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpCode, code int, data interface{}) {

	ctx.JSON(httpCode, gin.H{
		"code": code,
		"msg":  GetMsg(code),
		"data": data,
	})

}

func GetMsg(code int) string {
	str, ok := constants.MsgFlags[code]
	if ok {
		return str
	}
	return constants.MsgFlags[constants.ERROR]
}
