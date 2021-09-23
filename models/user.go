/*
 * @Author: lihuan
 * @Date: 2021-09-22 14:37:54
 * @LastEditTime: 2021-09-22 14:42:13
 * @Email: 17719495105@163.com
 */

package models


type User struct {
	UserName  string `json:"userName"`
	Password string `json:"password"`
}