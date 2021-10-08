/*
 * @Author: lihuan
 * @Date: 2021-09-23 14:03:14
 * @LastEditTime: 2021-10-08 17:28:40
 * @Email: 17719495105@163.com
 */

package dao

import (
	"go-lh/models"
	"go-lh/utils"

	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{DB: utils.GetDB()}
}

func (u *UserDao) Login(p *models.User) (*models.User, error) {
	var user models.User
	if err := u.DB.Where("user_name = ?", p.UserName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) Add(p *models.User) error {

	if err := u.DB.Create(p).Error; err != nil {
		return err
	}
	return nil

}

func (u *UserDao) GetUserInfo(id int) (*models.UserInfo, error) {

	var userInfo models.UserInfo

	if err := u.DB.Where("user_id = ?", id).First(&userInfo).Error; err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (u *UserDao) UpdateUserInfo(id int, uInfo *models.UserInfo) (*models.UserInfo, error) {
	var userInfo models.UserInfo

	if err := u.DB.Model(&userInfo).Where("user_id = ?", id).Updates(uInfo).Error; err != nil {
		return nil, err
	}

	return &userInfo, nil
}
