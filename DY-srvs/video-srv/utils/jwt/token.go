/*
 * @Date: 2023-01-22 17:38:05
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-26 10:55:45
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/jwt/token.go
 * @Description: 调用jwt产生Token
 */

package jwt

import (
	"simple-DY/DY-srvs/video-srv/global"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

/**
 * @description: 根据注册或者登录的id产生token
 * @param {int64} id
 * @return {string} token
 */
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

/**
 * @description: 解析Token
 * @param {string} token
 * @return {*}
 */
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

/**
 * @description: 解析Token，并与接收到的用户id作比较，如果一致则返回0
 * @param {string} token
 * @param {int64} userId
 * @return {int} StatusCode
 */
func GetAndJudgeIdByToken(token string, userId int64) int32 {
	// 没有携带Token信息
	if len(token) == 0 {
		zap.L().Error("没有携带Token信息！无法获取用户信息！")
		return 4
	}

	// 从Token中读取携带的id信息
	tokenId, err := ParseToken(strings.Fields(token)[1])
	if err != nil || tokenId.Id != strconv.FormatInt(userId, 10) {
		zap.L().Error("Token不正确！无法获取用户信息！")
		return 5
	}
	return 0
}
