/*
 * @Author: lihuan
 * @Date: 2021-09-22 13:30:04
 * @LastEditTime: 2021-09-23 15:42:37
 * @Email: 17719495105@163.com
 */

package controller

import (
	"go-lh/models"
	"go-lh/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {

}
func (u *UserController) Add(ctx *gin.Context)  {
	
	var user models.User
	
	ctx.BindJSON(&user)
	
	if err := service.UserSrv.Add(&user); err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"msg" : "OK",
			"data": err,
		})
		return
	}

	
	ctx.JSON(http.StatusOK,gin.H{
		"msg" : "OK",
		"data": nil,
	})
}