/*
 * @Author: lihuan
 * @Date: 2021-09-23 11:11:10
 * @LastEditTime: 2021-09-30 17:53:25
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
