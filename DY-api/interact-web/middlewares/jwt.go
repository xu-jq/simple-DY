/*
 * @Date: 2023-01-28 20:53:39
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-02 16:26:31
 * @FilePath: /simple-DY/DY-api/video-web/middlewares/jwt.go
 * @Description: JWT中间件
 */
package middlewares

import (
	"errors"
	"net/http"
	"simple-DY/DY-api/interact-web/global"
	"simple-DY/DY-api/interact-web/models"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		// token不知道怎么取，都是直接放在请求里面的，就Get和Post都试一下吧
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "请登录",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(strings.Fields(token)[1])
		if err != nil {
			zap.S().Info(err)
			if err == TokenExpired {
				if err == TokenExpired {
					c.JSON(http.StatusUnauthorized, map[string]string{
						"msg": "授权已过期",
					})
					c.Abort()
					return
				}
			}

			c.JSON(http.StatusUnauthorized, "未登录")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Set("TokenId", claims.Id)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	zap.S().Info(global.ServerConfig.JWTInfo.SigningKey)
	return &JWT{
		[]byte(global.ServerConfig.JWTInfo.SigningKey), //可以设置过期时间
	}
}

// 创建一个token
func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		zap.S().Info(err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func GenerateToken(id int64) string {
	zap.L().Info("开始产生Token...")

	// 设置Token过期时间
	expiresTime := time.Now().Unix() + 60*60*24
	zap.L().Info("Token将于" + time.Unix(expiresTime, 0).Format(global.ServerConfig.JWTInfo.SigningKey) + "过期")

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
	token, err := tokenClaims.SignedString([]byte(global.ServerConfig.JWTInfo.SigningKey))

	// 判断生成Token是否成功
	if err != nil {
		zap.L().Error("生成Token失败！错误信息：" + err.Error())
	} else {
		token = "Bearer " + token
		zap.L().Info("生成Token成功！")
	}
	return token
}
