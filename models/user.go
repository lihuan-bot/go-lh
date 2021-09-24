/*
 * @Author: lihuan
 * @Date: 2021-09-22 14:37:54
 * @LastEditTime: 2021-09-24 10:17:01
 * @Email: 17719495105@163.com
 */

package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	UserName  string `gorm:"varchar(20)" json:"userName"`
	Password string `gorm:"varchar(60)" json:"password"`
}