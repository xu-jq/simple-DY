/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-19 14:45:19
 * @FilePath: /simple-DY/DY-api/video-web/main.go
 * @Description:
 */
package main

import "simple-DY/DY-api/video-web/initialize"

func main() {
	r := initialize.Routers()
	r.Run(":8080")
}
