/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-28 22:35:12
 * @FilePath: /simple-DY/DY-api/video-web/api/base.go
 * @Description:
 */
package api

import (
	"net/http"
	"simple-DY/DY-api/video-web/global"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// func InitGRPC(port string) *grpc.ClientConn {
// 	addr := global.GlobalConfig.GRPC.Address + ":" + port
// 	size := global.GlobalConfig.GRPC.GRPCMsgSize.CommonMB
// 	// 上传文件需要将传输限制开大
// 	if port == "50001" {
// 		size = global.GlobalConfig.GRPC.GRPCMsgSize.LargeMB
// 	}
// 	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*size), grpc.MaxCallSendMsgSize(1024*1024*size)))
// 	if err != nil {
// 		zap.L().Error("初始化客户端GRPC失败！错误信息：" + err.Error())
// 	} else {
// 		zap.L().Info("初始化客户端GRPC成功！")
// 	}
// 	return conn
// }

func RemoveTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": RemoveTopStruct(errs.Translate(global.Trans)),
	})
}
