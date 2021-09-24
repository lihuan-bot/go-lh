/*
 * @Author: lihuan
 * @Date: 2021-09-23 14:03:14
 * @LastEditTime: 2021-09-24 10:14:40
 * @Email: 17719495105@163.com
 */

package dao

import (
	"go-lh/models"
	"go-lh/utils"
)

 type UserDao struct {

 }

 func (u *UserDao) Add(user *models.User) error {

	db := utils.NewDbEngine()
	
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
	
	
 }