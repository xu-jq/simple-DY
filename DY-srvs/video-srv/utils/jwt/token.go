/*
 * @Date: 2023-01-22 17:38:05
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 22:27:24
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/jwt/token.go
 * @Description: 调用jwt产生Token
 */

package jwt

import (
	"simple-DY/DY-srvs/video-srv/global"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

// 根据注册或者登录的id产生token
func GenerateToken(id int64) string {
	zap.L().Info("开始产生Token...")

	// 设置Token过期时间
	expiresTime := time.Now().Unix() + global.GlobalConfig.JWT.TokenExpiresTime
	zap.L().Info("Token将于" + time.Unix(expiresTime, 0).Format(global.GlobalConfig.Time.TimeFormat) + "过期")

	// 声明
	claims := jwt.StandardClaims{
		ExpiresAt: expiresTime,
		Id:        strconv.FormatInt(id, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "simple-dy",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}

	// 生成Token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(global.GlobalConfig.JWT.Secret))

	// 判断生成Token是否成功
	if err != nil {
		zap.L().Error("生成Token失败！错误信息：" + err.Error())
	} else {
		token = "Bearer " + token
		zap.L().Info("生成Token成功！")
	}
	return token
}

// 解析token
func ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(global.GlobalConfig.JWT.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			zap.L().Info("Token解析成功！")
			return claim, nil
		}
	}
	zap.L().Error("Token解析失败！错误信息：" + err.Error())
	return nil, err
}
