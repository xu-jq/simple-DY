/*
 * @Date: 2023-02-05 19:22:56
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-02-05 19:23:16
 * @FilePath: /simple-DY/DY-srvs/video-srv/handler/base.go
 * @Description: Videoserver
 */
package handler

import (
	pb "simple-DY/DY-srvs/video-srv/proto"
)

type Videoserver struct {
	pb.UnimplementedVideoServiceServer
}
