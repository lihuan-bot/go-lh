/*
 * @Author: lihuan
 * @Date: 2021-09-22 13:22:13
 * @LastEditTime: 2021-10-08 17:27:08
 * @Email: 17719495105@163.com
 */
package router

import (
	"go-lh/controller"
	"go-lh/middleware"
	"go-lh/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *utils.Config) *gin.Engine {
	r := gin.Default()

	gin.SetMode(cfg.AppMode)

	v1Group := r.Group("v1")

	v1Group.POST("/user/login", controller.UserCtl.Login)

	// 使用 token 中间
	v1Group.Use(middleware.JWT())

	{

		v1Group.POST("/user/add", controller.UserCtl.Add)

		v1Group.GET("/user/userInfo/:id", controller.UserCtl.GetUserInfo)
		v1Group.PUT("/user/userInfo/:id", controller.UserCtl.UpdateUserInfo)

	}
	return r
}
