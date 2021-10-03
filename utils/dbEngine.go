/*
 * @Author: lihuan
 * @Date: 2021-09-23 08:49:14
 * @LastEditTime: 2021-09-30 17:56:01
 * @Email: 17719495105@163.com
 */
package utils

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func DBInit(cfg *Config) error {
	c := cfg.Database
	// 拼接链接串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%v&loc=%s", c.User, c.Password, c.Host, c.Port, c.DbName, c.Charset, c.ParseTime, c.Loc)

	engine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = engine

	return err

}
