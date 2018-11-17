/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50621
Source Host           : localhost:3306
Source Database       : golang

Target Server Type    : MYSQL
Target Server Version : 50621
File Encoding         : 65001

Date: 2018-11-17 17:47:25
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for room
-- ----------------------------
DROP TABLE IF EXISTS `room`;
CREATE TABLE `room` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `room_id` bigint(20) NOT NULL,
  `room_name` varchar(64) NOT NULL,
  `desc` varchar(255) NOT NULL,
  `online` int(10) unsigned DEFAULT '0',
  `status` int(10) unsigned NOT NULL DEFAULT '1',
  `cap` int(11) NOT NULL DEFAULT '500',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_room_id` (`room_id`),
  UNIQUE KEY `idx_room_name` (`room_name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of room
-- ----------------------------
INSERT INTO `room` VALUES ('1', '322', '午夜情场', '你寂寞吗', '0', '1', '500', '2018-11-17 17:39:00', '2018-11-17 17:43:53');
INSERT INTO `room` VALUES ('2', '323', '午夜码农', '一起来学习golang', '0', '1', '500', '2018-11-17 17:44:31', '2018-11-17 17:44:42');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `username` varchar(64) NOT NULL DEFAULT '',
  `nickname` varchar(64) NOT NULL DEFAULT '0',
  `password` varchar(64) NOT NULL,
  `sex` tinyint(4) NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`),
  UNIQUE KEY `idx_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('44', '12', 'aaaa', 'aadd22', '', '0', '2018-11-17 10:33:14', '2018-11-17 10:33:23');
INSERT INTO `user` VALUES ('45', '222964304830267393', 'carry', '五连败', '', '0', '2018-11-17 11:51:21', '2018-11-17 11:51:21');
INSERT INTO `user` VALUES ('48', '222964959611453441', 'carry01', '五连败', '0c1f76cd0ea011312f75ccdf8cf83b86', '0', '2018-11-17 11:57:52', '2018-11-17 11:57:52');
INSERT INTO `user` VALUES ('49', '222981981489594369', 'admin', 'xuliangwei', 'bb7180abeb377d0e51f4ea8c00a6126b', '0', '2018-11-17 14:46:57', '2018-11-17 14:46:57');
INSERT INTO `user` VALUES ('55', '222982053396742145', 'huas228', 'shalo', 'bb7180abeb377d0e51f4ea8c00a6126b', '0', '2018-11-17 14:47:40', '2018-11-17 14:47:40');
INSERT INTO `user` VALUES ('64', '222982394645315585', 'admin45444', 'xuliangwei', 'bb7180abeb377d0e51f4ea8c00a6126b', '0', '2018-11-17 14:51:04', '2018-11-17 14:51:04');
INSERT INTO `user` VALUES ('66', '222982813303963649', 'admin123', 'xuliangwei', 'bb7180abeb377d0e51f4ea8c00a6126b', '0', '2018-11-17 14:55:13', '2018-11-17 14:55:13');
INSERT INTO `user` VALUES ('67', '222982884909121537', 'kuangqianfei', '', '24a0f7cb1f1f956d006b426a5124cbab', '0', '2018-11-17 14:55:56', '2018-11-17 14:55:56');
INSERT INTO `user` VALUES ('68', '222982914638348289', 'username', 'nickmae', 'bb7180abeb377d0e51f4ea8c00a6126b', '0', '2018-11-17 14:56:14', '2018-11-17 14:56:14');
INSERT INTO `user` VALUES ('69', '222983454613045249', 'admin2', 'xuliangwei', '906c60f3a2264a7f268149522f761695', '0', '2018-11-17 15:01:36', '2018-11-17 15:01:36');
INSERT INTO `user` VALUES ('71', '222987883814846465', 'admin2018', 'admin2018', 'a2b47b1790593091052fd7ff1a7c398f', '0', '2018-11-17 15:45:36', '2018-11-17 15:45:36');
