/*
 * @Date: 2023-01-19 18:39:07
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-20 19:31:55
 * @FilePath: /simple-DY/DY-srvs/video-srv/config/video.sql
 * @Description: 数据库初始SQL操作
 */

-- CREATE USER 'dymysql'@'%' IDENTIFIED BY 'gxnw21XxRhY';
-- ALTER USER 'dymysql'@'%' IDENTIFIED WITH mysql_native_password BY 'gxnw21XxRhY';
-- GRANT ALL PRIVILEGES ON `simpledy`.* TO `dymysql`@`%` WITH GRANT OPTION;
-- -- REVOKE ALL PRIVILEGES, GRANT OPTION FROM 'dymysql';
-- FLUSH PRIVILEGES;
-- SHOW GRANTS FOR 'dymysql'@'%';

SET NAMES utf8mb4;
USE simpledy;
-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '用户id，自增主键',
  `name` VARCHAR(255) NOT NULL COMMENT '用户名',
  `password` VARCHAR(255) NOT NULL COMMENT '用户密码',
  PRIMARY KEY (`id`),
  KEY `name_password_idx` (`name`,`password`) USING BTREE
) DEFAULT CHARSET=utf8 COMMENT='用户表';

INSERT INTO users (name,password) VALUES ('ZhangZhao','ZhangZhaoPassword');

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键，视频唯一id',
  `author_id` BIGINT(20) NOT NULL COMMENT '视频作者id',
  `file_name` VARCHAR(255) NOT NULL COMMENT '文件命名',
  `video_suffix` CHAR(10) NOT NULL COMMENT '视频后缀',
  `publish_time` datetime NOT NULL COMMENT '发布时间',
  `title` VARCHAR(255) DEFAULT NULL COMMENT '视频标题',
  PRIMARY KEY (`id`),
  KEY `time` (`publish_time`) USING BTREE,
  KEY `author` (`author_id`) USING BTREE
) DEFAULT CHARSET=utf8 COMMENT='视频表';

INSERT INTO videos (author_id, file_name, video_suffix, publish_time, title) VALUES (1,'example1','.mp4',now(),'第一个视频');
INSERT INTO videos (author_id, file_name, video_suffix, publish_time, title) VALUES (1,'example2','.mp4',now(),'第二个视频');
INSERT INTO videos (author_id, file_name, video_suffix, publish_time, title) VALUES (1,'example3','.mp4',now(),'第三个视频');