/*
 * @Date: 2023-01-29 10:11:07
 * @LastEditors: wang hui, zhang zhao
 * @LastEditTime: 2023-01-29 10:16:03
 * @FilePath: /simple-DY/DY-api/video-web/models/jwt.go
 * @Description: JWT结构体
 */

package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}
