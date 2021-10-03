/*
 * @Author: lihuan
 * @Date: 2021-09-30 10:31:04
 * @LastEditTime: 2021-09-30 17:56:15
 * @Email: 17719495105@163.com
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string `json:"username"`
	ID       int    `json:"id"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func GenerateToken(id int, username, password string) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(Cfg.Token.ExpireTime) * time.Hour)
	fmt.Println(time.Duration(Cfg.Token.ExpireTime), 9, Cfg.Token.ExpireTime, Cfg.Token.Secret)
	claims := Claims{
		username,
		id,
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-lh",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(Cfg.Token.Secret))

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Cfg.Token.Secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
