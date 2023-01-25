/*
 * @Date: 2023-01-19 18:39:07
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-25 23:10:01
 * @FilePath: /simple-DY/db.sql
 * @Description: 数据库初始SQL操作
 */

-- CREATE USER 'dymysql'@'%' IDENTIFIED BY 'gxnw21XxRhY';
-- ALTER USER 'dymysql'@'%' IDENTIFIED WITH mysql_native_password BY 'gxnw21XxRhY';
-- GRANT ALL PRIVILEGES ON `simpledy`.* TO `dymysql`@`%` WITH GRANT OPTION;
-- -- REVOKE ALL PRIVILEGES, GRANT OPTION FROM 'dymysql';
-- FLUSH PRIVILEGES;
-- SHOW GRANTS FOR 'dymysql'@'%';

-- USE simpledy;


-- comments definition

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


-- follow  definition

DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `follower_id` bigint NOT NULL COMMENT '关注的用户',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_follower_idx` (`user_id`,`follower_id`) USING BTREE COMMENT '关注用户间联合索引',
  KEY `follower_idx` (`follower_id`) USING BTREE COMMENT '关注用户的索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='关注表';


-- likes  definition

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


-- messages definition

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


-- users  definition

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户id，自增主键',
  `name` varchar(255) BINARY NOT NULL COMMENT '用户名',
  `password` varchar(255) BINARY NOT NULL COMMENT '用户密码',
  PRIMARY KEY (`id`),
  KEY `name_password_idx` (`name`,`password`) USING BTREE COMMENT '用户名和密码的联合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='用户表';

INSERT INTO users (id, name, password) VALUES(1, 'test1', '4a3252a5edf8fcaa8bde0bfcce79560d');
INSERT INTO users (id, name, password) VALUES(2, 'test2', '80660e29103d525b694f45e34e23f498');
INSERT INTO users (id, name, password) VALUES(3, 'test3', 'ed05155fbf4f7a6373bc7c344be065bd');


-- videos definition

DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频唯一id，自增主键',
  `author_id` bigint NOT NULL COMMENT '视频作者id',
  `file_name` varchar(255) NOT NULL COMMENT '文件命名',
  `publish_time` bigint NOT NULL COMMENT '发布时间',
  `title` varchar(255) DEFAULT NULL COMMENT '视频标题',
  PRIMARY KEY (`id`),
  KEY `time_idx` (`publish_time`) USING BTREE COMMENT '发布时间索引',
  KEY `author_idx` (`author_id`) USING BTREE COMMENT '作者id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='视频表';

INSERT INTO videos (id, author_id, file_name, publish_time, title) VALUES(1, 1, 'f5316905-5a72-4380-8979-6fbc178c1ba9', 1674477443, '1粉色');
INSERT INTO videos (id, author_id, file_name, publish_time, title) VALUES(2, 1, '444be6f6-92ce-4cab-a506-c268777a80d0', 1674477471, '2粉红色');
INSERT INTO videos (id, author_id, file_name, publish_time, title) VALUES(3, 2, 'd6f03720-f0f6-47cd-8b88-01fbf827e38a', 1674477559, '3粉色');
INSERT INTO videos (id, author_id, file_name, publish_time, title) VALUES(4, 2, 'e97e550f-5b41-4603-b737-8e72b80a74e1', 1674477585, '4粉色');
INSERT INTO videos (id, author_id, file_name, publish_time, title) VALUES(5, 2, '51f50227-d61a-435f-9379-44e5c0220689', 1674477606, '5粉色');
INSERT INTO videos (id, author_id, file_name, publish_time, title) VALUES(6, 3, '64647d28-3853-4414-b6ad-8f5b43fe9c5c', 1674477656, '6紫色');
INSERT INTO videos (id, author_id, file_name, publish_time, title) VALUES(7, 3, '360f2aeb-072c-47ce-bc2c-003e0f285df2', 1674477676, '7紫色');
INSERT INTO videos (id, author_id, file_name, publish_time, title) VALUES(8, 3, 'fc222697-e128-4f14-abb1-da8d04bdb44b', 1674477693, '8紫色');
INSERT INTO videos (id, author_id, file_name, publish_time, title) VALUES(9, 3, '675106e1-0f1a-4055-bd22-6995d7cc3417', 1674477706, '9蓝色');
INSERT INTO videos (id, author_id, file_name, publish_time, title) VALUES(10, 3, 'beec5a6f-3ea1-4e18-bd5a-196865bc3f3c', 1674477721, '10浅蓝色');
