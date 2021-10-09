/*
 * @Author: lihuan
 * @Date: 2021-09-24 09:40:24
 * @LastEditTime: 2021-10-09 11:14:58
 * @Email: 17719495105@163.com
 */
package constants

const (
	SUCCESS               = 2000
	NOT_FOUND_USER        = 2001
	PASSWORD_ERROR        = 2002
	GENERATE_TOKEN_ERROR  = 2003
	INVALID_TOKEN_ERROR   = 2004
	TOKEN_TIMEOUT_ERROR   = 2005
	TOKEN_FAIL            = 2006
	UNAUTHORIZED_ERROR    = 2007
	GET_USERINFO_ERROR    = 2008
	UPDATE_USERINFO_ERROR = 2009
	PARAMS_ISVALIID       = 2010
	PARAMS_RULES_ERROR    = 2011
	RECORD_NOT_FOUND      = 3000
	ERROR                 = 5000
)

var MsgFlags = map[int]string{
	SUCCESS:               "成功",
	NOT_FOUND_USER:        "用户名不存在",
	PASSWORD_ERROR:        "密码不正确",
	GENERATE_TOKEN_ERROR:  "Token 生成失败",
	INVALID_TOKEN_ERROR:   "无效的 Token",
	TOKEN_TIMEOUT_ERROR:   "token过期",
	UNAUTHORIZED_ERROR:    "没有权限",
	GET_USERINFO_ERROR:    "获取用户信息失败",
	UPDATE_USERINFO_ERROR: "修改用户信息失败",
	TOKEN_FAIL:            "解析token失败",
	PARAMS_ISVALIID:       "参数是无效的",
	PARAMS_RULES_ERROR:    "参数规则校验错误",
	RECORD_NOT_FOUND:      "未匹配到该记录",
	ERROR:                 "未知错误",
}
