/*
 * @Author: lihuan
 * @Date: 2021-09-22 14:37:54
 * @LastEditTime: 2021-09-30 17:55:30
 * @Email: 17719495105@163.com
 */

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"size:20;unique;comment:用户名" json:"userName"`
	Password string `gorm:"size:60;comment:密码" json:"password"`
}
