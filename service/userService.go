/*
 * @Author: lihuan
 * @Date: 2021-09-23 11:11:10
 * @LastEditTime: 2021-09-24 09:01:43
 * @Email: 17719495105@163.com
 */
package service

import (
	"go-lh/dao"
	"go-lh/models"
)

type UserService struct {
	
}
func (u *UserService) Add(user *models.User) error{
	
	if err := dao.Userdao.Add(user); err != nil {
		return err
	}
	
	return nil
}

