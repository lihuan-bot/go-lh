/*
 * @Author: lihuan
 * @Date: 2021-09-22 13:30:04
 * @LastEditTime: 2021-10-09 13:36:11
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
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
}

// 用户登录
func (u *UserController) Login(ctx *gin.Context) {
	var user models.User
	// 检验参数
	if err := ctx.ShouldBindWith(&user, binding.JSON); err != nil {
		errs := err.(validator.ValidationErrors)
		utils.Response(ctx, http.StatusBadRequest, constants.PARAMS_RULES_ERROR, utils.ErrorTranslate(errs))
		return
	}

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

	utils.Response(ctx, http.StatusOK, constants.SUCCESS, gin.H{"username": us.UserName, "id": us.ID, "token": token})
}

// 新增用户
func (u *UserController) Add(ctx *gin.Context) {

	var user models.User

	// ctx.BindJSON(&user)
	// 检验参数
	if err := ctx.ShouldBindWith(&user, binding.JSON); err != nil {
		errs := err.(validator.ValidationErrors)
		utils.Response(ctx, http.StatusBadRequest, constants.PARAMS_RULES_ERROR, utils.ErrorTranslate(errs))
		return
	}

	if err := service.UserSrv.Add(&user); err != nil {
		utils.Response(ctx, http.StatusOK, constants.ERROR, nil)
		return
	}

	utils.Response(ctx, http.StatusOK, constants.SUCCESS, nil)

}

// 获取用户信息
func (u *UserController) GetUserInfo(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	userInfo, err := service.UserSrv.GetUserInfo(id)

	if err != nil {
		utils.Response(ctx, http.StatusOK, constants.GET_USERINFO_ERROR, nil)
		return
	}

	utils.Response(ctx, http.StatusOK, constants.SUCCESS, userInfo)

}

// 更新用户信息
func (u *UserController) UpdateUserInfo(ctx *gin.Context) {

	var userInfo models.UserInfo

	// ctx.BindJSON(&userInfo)
	if err := ctx.ShouldBindWith(&userInfo, binding.JSON); err != nil {
		errs := err.(validator.ValidationErrors)
		utils.Response(ctx, http.StatusBadRequest, constants.PARAMS_RULES_ERROR, utils.ErrorTranslate(errs))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	uInfo, err := service.UserSrv.UpdateUserInfo(id, &userInfo)
	if err != nil {
		utils.Response(ctx, http.StatusOK, constants.UPDATE_USERINFO_ERROR, nil)
		return
	}
	utils.Response(ctx, http.StatusOK, constants.SUCCESS, uInfo)

}
