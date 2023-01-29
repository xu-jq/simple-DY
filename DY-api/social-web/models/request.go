package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID          int32
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}
