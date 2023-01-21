/*
 * @Date: 2023-01-19 18:39:07
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-21 17:31:35
 * @FilePath: /simple-DY/DY-srvs/video-srv/config/video.sql
 * @Description: 数据库初始SQL操作
 */

-- CREATE USER 'dymysql'@'%' IDENTIFIED BY 'gxnw21XxRhY';
-- ALTER USER 'dymysql'@'%' IDENTIFIED WITH mysql_native_password BY 'gxnw21XxRhY';
-- GRANT ALL PRIVILEGES ON `simpledy`.* TO `dymysql`@`%` WITH GRANT OPTION;
-- -- REVOKE ALL PRIVILEGES, GRANT OPTION FROM 'dymysql';
-- FLUSH PRIVILEGES;
-- SHOW GRANTS FOR 'dymysql'@'%';

-- USE simpledy;


-- simpledy.comments definition

DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论id，自增主键',
  `user_id` bigint NOT NULL COMMENT '评论发布用户id',
  `video_id` bigint NOT NULL COMMENT '被评论视频的id',
  `comment_text` varchar(255) NOT NULL COMMENT '评论内容',
  `create_time` datetime NOT NULL COMMENT '评论发布时间',
  PRIMARY KEY (`id`),
  KEY `video_idx` (`video_id`) USING BTREE COMMENT '视频id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评论表';


-- simpledy.follows definition

DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `follower_id` bigint NOT NULL COMMENT '关注的用户',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_follower_idx` (`user_id`,`follower_id`) USING BTREE COMMENT '关注用户间联合索引',
  KEY `follower_idx` (`follower_id`) USING BTREE COMMENT '关注用户的索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='关注表';


-- simpledy.likes definition

DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint NOT NULL COMMENT '点赞用户id',
  `video_id` bigint NOT NULL COMMENT '被点赞的视频id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_video_idx` (`user_id`,`video_id`) USING BTREE COMMENT '用户和点赞视频联合索引',
  KEY `user_idx` (`user_id`) USING BTREE COMMENT '用户id索引',
  KEY `video_idx` (`video_id`) USING BTREE COMMENT '视频id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='点赞表';


-- simpledy.messages definition

DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint NOT NULL COMMENT '发送消息的用户id',
  `to_user_id` bigint NOT NULL COMMENT '接收消息的用户id',
  `sent_time` datetime NOT NULL COMMENT '消息发送时间',
  `content` varchar(255) NOT NULL COMMENT '消息内容',
  PRIMARY KEY (`id`),
  KEY `user_idx` (`user_id`) USING BTREE COMMENT '发送用户id索引',
  KEY `to_user_idx` (`to_user_id`) USING BTREE COMMENT '接收用户id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='消息表';


-- simpledy.users definition

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户id，自增主键',
  `name` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  PRIMARY KEY (`id`),
  KEY `name_password_idx` (`name`,`password`) USING BTREE COMMENT '用户名和密码的联合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='用户表';

INSERT INTO users (name,password) VALUES ('ZhangZhao','ZhangZhaoPassword');


-- simpledy.videos definition

DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频唯一id，自增主键',
  `author_id` bigint NOT NULL COMMENT '视频作者id',
  `file_name` varchar(255) NOT NULL COMMENT '文件命名',
  `video_suffix` char(10) NOT NULL COMMENT '视频后缀',
  `publish_time` datetime NOT NULL COMMENT '发布时间',
  `title` varchar(255) DEFAULT NULL COMMENT '视频标题',
  PRIMARY KEY (`id`),
  KEY `time_idx` (`publish_time`) USING BTREE COMMENT '发布时间索引',
  KEY `author_idx` (`author_id`) USING BTREE COMMENT '作者id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='视频表';

INSERT INTO videos (author_id, file_name, video_suffix, publish_time, title) VALUES (1,'example1','.mp4',now(),'第一个视频');
INSERT INTO videos (author_id, file_name, video_suffix, publish_time, title) VALUES (1,'example2','.mp4',now(),'第二个视频');
INSERT INTO videos (author_id, file_name, video_suffix, publish_time, title) VALUES (1,'example3','.mp4',now(),'第三个视频');
