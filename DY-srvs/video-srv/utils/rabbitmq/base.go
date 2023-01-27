/*
 * @Date: 2023-01-26 22:44:52
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-27 10:19:43
 * @FilePath: /simple-DY/DY-srvs/video-srv/utils/rabbitmq/base.go
 * @Description: 消息结构体
 */
package rabbitmq

type Message struct {
	VideoStaticFileName string `json:"video_static_file_name"`
	VideoOSSFileName    string `json:"video_oss_file_name"`
	ImageStaticFileName string `json:"image_static_file_name"`
	ImageOSSFileName    string `json:"image_oss_file_name"`
	AuthorId            int64  `json:"author_id"`
	FileName            string `json:"file_name"`
	Time                int64  `json:"time"`
	Title               string `json:"title"`
}
