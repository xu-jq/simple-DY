/*
 * @Date: 2023-01-22 10:36:26
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-22 15:09:56
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/md5salt/md5.go
 * @Description: md5 加密函数
 */
package md5salt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5V(str string, salt string, iteration int) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先传入盐值，之前因为顺序错了卡了很久
	h.Write(b)
	var res []byte
	res = h.Sum(nil)
	for i := 0; i < iteration-1; i++ {
		h.Reset()
		h.Write(res)
		res = h.Sum(nil)
	}
	return hex.EncodeToString(res)
}
