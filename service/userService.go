/*
 * @Author: lihuan
 * @Date: 2021-09-23 11:11:10
 * @LastEditTime: 2021-10-08 17:17:16
 * @Email: 17719495105@163.com
 */
package service

import (
	"go-lh/dao"
	"go-lh/models"
	"go-lh/utils"
)

type UserService struct {
}

func (u *UserService) Login(p *models.User) (*models.User, error) {

	ud := dao.NewUserDao()

	user, err := ud.Login(p)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *UserService) Add(p *models.User) error {

	ud := dao.NewUserDao()

	p.Password = utils.EncodeMD5(p.Password)

	if err := ud.Add(p); err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetUserInfo(id int) (*models.UserInfo, error) {
	ud := dao.NewUserDao()
	userInfo, err := ud.GetUserInfo(id)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (u *UserService) UpdateUserInfo(id int, uInfo *models.UserInfo) (*models.UserInfo, error) {
	ud := dao.NewUserDao()

	userInfo, err := ud.UpdateUserInfo(id, uInfo)

	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
