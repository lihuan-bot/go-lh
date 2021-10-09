/*
 * @Author: lihuan
 * @Date: 2021-09-22 14:37:54
 * @LastEditTime: 2021-10-09 22:03:08
 * @Email: 17719495105@163.com
 */

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string   `gorm:"size:30;unique;comment:用户名" json:"userName" binding:"required,max=30,min=1" label:"用户名"`
	Password string   `gorm:"size:60;comment:密码" json:"password" binding:"required,max=60,min=6" label:"密码"`
	UserInfo UserInfo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserInfo struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"-"`
	Avater  string `gorm:"size:300;comment:头像" json:"avater"`
	Gender  string `gorm:"size:4;default:'男';comment:性别" json:"gender"`
	Address string `gorm:"size:100;comment:地址" json:"address"`
	Hobby   string `gorm:"size:100;comment:爱好" json:"hobby"`
	UserID  int    `json:"-"`
}
