/*
 * @Author: lihuan
 * @Date: 2021-09-23 08:49:14
 * @LastEditTime: 2021-09-23 17:27:38
 * @Email: 17719495105@163.com
 */
package utils

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

  var db *gorm.DB
  
  type DBEngine struct {
	  *gorm.DB
  }

  func NewDbEngine () *DBEngine{
	return &DBEngine{DB:  db}
  }
  
  func DBInit(cfg *Config) error {
	c := cfg.Database
	// 拼接链接串
	dsn := fmt.Sprintf(c.User +  ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.DbName + "?charset=" + c.Charset + "&parseTime" + c.Charset + "&loc=" + c.Loc)
	engine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	db	= engine 

	return err
	
  }