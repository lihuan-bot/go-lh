/*
 * @Author: lihuan
 * @Date: 2021-09-22 13:30:04
 * @LastEditTime: 2021-10-02 09:58:08
 * @Email: 17719495105@163.com
 */

package controller

import (
	"fmt"
	"go-lh/constants"
	"go-lh/models"
	"go-lh/service"
	"go-lh/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 用户登录
func (u *UserController) Login(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)

	us, err := service.UserSrv.Login(&user)
	// 用户不存在
	if err != nil {
		utils.Response(ctx, http.StatusOK, constants.NOT_FOUND_USER, nil)
		fmt.Println(err)
		return
	}
	// 密码不正确
	if us.Password != utils.EncodeMD5(user.Password) {
		utils.Response(ctx, http.StatusOK, constants.PASSWORD_ERROR, nil)
		fmt.Println(err)
		return
	}
	// 生成 token
	token, err := utils.GenerateToken(int(us.ID), us.UserName, user.Password)
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, constants.GENERATE_TOKEN_ERROR, nil)
		fmt.Println(err)
		return
	}

	utils.Response(ctx, http.StatusOK, constants.SUCCESS, gin.H{"username": us.UserName, "token": token})
}

// 新增用户
func (u *UserController) Add(ctx *gin.Context) {

	var user models.User

	ctx.BindJSON(&user)

	if err := service.UserSrv.Add(&user); err != nil {
		utils.Response(ctx, http.StatusOK, constants.ERROR, err)
		return
	}

	utils.Response(ctx, http.StatusOK, constants.SUCCESS, nil)

}
