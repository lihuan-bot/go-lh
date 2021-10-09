/*
 * @Author: lihuan
 * @Date: 2021-09-22 11:34:27
 * @LastEditTime: 2021-10-09 09:47:39
 * @Email: 17719495105@163.com
 */

package main

import (
	"fmt"
	"go-lh/models"
	"go-lh/router"
	"go-lh/utils"
)

func main() {

	// 初始化配置文件
	cfg, err := utils.ParseConfig("./config/app.json")

	if err != nil {
		panic(err.Error())
	}
	// 初始化数据库
	err = utils.DBInit(cfg)
	if err != nil {
		fmt.Println(err)
	}

	db := utils.GetDB()
	db.AutoMigrate(&models.User{}, &models.UserInfo{})

	// 注册路由
	r := router.SetupRouter(cfg)
	addr := fmt.Sprintf(cfg.AppHost + ":" + cfg.AppPort)

	r.Run(addr)
}
