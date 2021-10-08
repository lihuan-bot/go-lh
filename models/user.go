/*
 * @Author: lihuan
 * @Date: 2021-09-22 14:37:54
 * @LastEditTime: 2021-10-08 16:30:44
 * @Email: 17719495105@163.com
 */

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string   `gorm:"size:20;unique;comment:用户名" json:"userName"`
	Password string   `gorm:"size:60;comment:密码" json:"password"`
	UserInfo UserInfo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserInfo struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"-"`
	Avater  string `gorm:"size:300;comment:头像" json:"avater"`
	Gender  string `gorm:"size:4;default:'男';check:Gender in ('男','女');comment:性别" json:"gender"`
	Address string `gorm:"size:100;comment:地址" json:"address"`
	Hobby   string `gorm:"size:100;comment:爱好" json:"hobby"`
	UserID  int    `json:"-"`
}
