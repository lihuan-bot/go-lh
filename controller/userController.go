/*
 * @Author: lihuan
 * @Date: 2021-09-22 13:30:04
 * @LastEditTime: 2021-09-24 09:55:52
 * @Email: 17719495105@163.com
 */

package controller

import (
	"go-lh/constants"
	"go-lh/models"
	"go-lh/service"
	"go-lh/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {

}
func (u *UserController) Add(ctx *gin.Context)  {

	var user models.User
	
	ctx.BindJSON(&user)
	
	if err := service.UserSrv.Add(&user); err != nil {
		utils.Response(ctx,http.StatusOK,constants.ERROR,gin.H{"err": err})
		return
	}

	 utils.Response(ctx,http.StatusOK,constants.SUCCESS,gin.H{})
	
}