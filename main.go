/*
 * @Author: lihuan
 * @Date: 2021-09-22 11:34:27
 * @LastEditTime: 2021-09-23 16:17:14
 * @Email: 17719495105@163.com
 */

package main

import (
	"fmt"
	"go-lh/models"
	"go-lh/router"
	"go-lh/utils"
)



func main()  {
	 
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
	 e := utils.NewDbEngine()
	 e.AutoMigrate(&models.User{})
	 fmt.Println(1111111,e)
	// 注册路由
	r := router.SetupRouter()
	addr := fmt.Sprintf(cfg.AppHost+ ":" + cfg.AppPort)
	
	r.Run(addr)
}