/*
 * @Author: lihuan
 * @Date: 2021-09-22 13:22:13
 * @LastEditTime: 2021-09-24 10:19:54
 * @Email: 17719495105@163.com
 */
package router

import (
	"go-lh/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	


	v1Group := r.Group("v1")
	{
		
		v1Group.POST("/user", controller.UserCtl.Add)

	}
	return r
}