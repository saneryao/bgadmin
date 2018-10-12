/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : saneryao_goadmin

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 12/10/2018 10:59:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for 321_code
-- ----------------------------
DROP TABLE IF EXISTS `321_code`;
CREATE TABLE `321_code`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT 0,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `create_time` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_code
-- ----------------------------
INSERT INTO `321_code` VALUES (1, 41, '0lEHGWmsysP92EMim1M6faN7QBl4t2KaRrp6hUFtk4fGOUffEviAxrkfinMVLPHi', 1537894195);
INSERT INTO `321_code` VALUES (2, 42, 'YUbQyK5byDpMQ2pJq6SXwGeEfB5aBk9DV6gYESwSWgp8DHupzrHUBkiylpR74WQ5', 1537894642);
INSERT INTO `321_code` VALUES (3, 43, 'dBYNDmhnvcEtdC3md9NTwXfhvQzM0uR4sMzkionNcNSuFNBRJDarZ1IeKOqRBBfi', 1537894732);
INSERT INTO `321_code` VALUES (4, 44, 'W3xaWYzkzdolZwvWhAHf2cKXTIKNodVXFPdkD72V9nZYAh7jr6hwdHkwm1KoqlLc', 1537894911);

-- ----------------------------
-- Table structure for 321_menu
-- ----------------------------
DROP TABLE IF EXISTS `321_menu`;
CREATE TABLE `321_menu`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) NOT NULL DEFAULT 0,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `state` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_menu
-- ----------------------------
INSERT INTO `321_menu` VALUES (7, 0, '系统首页', 'fa-home', '', 1);
INSERT INTO `321_menu` VALUES (9, 0, '用户和权限', 'fa-users', '', 1);
INSERT INTO `321_menu` VALUES (17, 9, '菜单', '', '/#page/menu', 1);
INSERT INTO `321_menu` VALUES (15, 9, '角色', '', '/#page/role', 1);
INSERT INTO `321_menu` VALUES (11, 0, '更多示例', 'fa-info-circle', '', 0);
INSERT INTO `321_menu` VALUES (14, 9, '用户', '', '/#page/user', 1);
INSERT INTO `321_menu` VALUES (18, 11, '页面', '', '', 1);
INSERT INTO `321_menu` VALUES (19, 11, '排版', '', '', 1);
INSERT INTO `321_menu` VALUES (20, 11, '按钮', '', '', 1);
INSERT INTO `321_menu` VALUES (21, 11, '图标', '', '', 1);

-- ----------------------------
-- Table structure for 321_profile
-- ----------------------------
DROP TABLE IF EXISTS `321_profile`;
CREATE TABLE `321_profile`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sex` int(11) NULL DEFAULT NULL,
  `birth` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `nick` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `mobile` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `tel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_profile
-- ----------------------------
INSERT INTO `321_profile` VALUES (1, 1, NULL, 'admin*', 'admin@saneryao.com', '1301106553', NULL, '');
INSERT INTO `321_profile` VALUES (2, 1, '', 'test', 'test@saneryao.com', '', '', '');
INSERT INTO `321_profile` VALUES (3, 1, '', 'youyou', 'youyou@saneryao.com', '', '', '');
INSERT INTO `321_profile` VALUES (4, 1, '', 'demo', 'demo@saneryao.com', '', '', '');
INSERT INTO `321_profile` VALUES (5, 1, '', 'root', 'root@saneryao.com', '', '', '');
INSERT INTO `321_profile` VALUES (6, 1, '', 'saneryao', 'saneryao@saneryao.com', '', '', '');
INSERT INTO `321_profile` VALUES (29, 1, '2018-10-04', 'haha', 'haha@saneryao.com', '', '', 'Beijing');
INSERT INTO `321_profile` VALUES (39, 1, '2018-10-03', '...', 'abcd@saneryao.com', '', '', 'Beijing');
INSERT INTO `321_profile` VALUES (40, 1, '2018-09-26', 'aaaa...', 'aaaa@saneryao.com', '', '', 'Beijing');
INSERT INTO `321_profile` VALUES (44, 0, '', '', '465245445@qq.com', '', '', '');

-- ----------------------------
-- Table structure for 321_resource
-- ----------------------------
DROP TABLE IF EXISTS `321_resource`;
CREATE TABLE `321_resource`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `state` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for 321_role
-- ----------------------------
DROP TABLE IF EXISTS `321_role`;
CREATE TABLE `321_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `state` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_role
-- ----------------------------
INSERT INTO `321_role` VALUES (1, 'admin', 'admin', 1);
INSERT INTO `321_role` VALUES (3, 'test', 'test...', 0);
INSERT INTO `321_role` VALUES (4, 'abcd', '...', 1);
INSERT INTO `321_role` VALUES (6, 'aaaa', '', 0);

-- ----------------------------
-- Table structure for 321_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `321_role_menu`;
CREATE TABLE `321_role_menu`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` bigint(20) NOT NULL DEFAULT 0,
  `menu_id` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Fixed;

-- ----------------------------
-- Table structure for 321_user
-- ----------------------------
DROP TABLE IF EXISTS `321_user`;
CREATE TABLE `321_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `pwd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `state` int(11) NOT NULL DEFAULT 0,
  `profile_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `profile_id`(`profile_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_user
-- ----------------------------
INSERT INTO `321_user` VALUES (1, 'admin', 'AQpyZkJkNTZ0aTJTiMp2eLHYq5Ex+roSKU2p58M2sKp2dt5fmUBi6x8+MJY=', 1, 1);
INSERT INTO `321_user` VALUES (2, 'test', 'AQpoUGprU0J2UU9WCcWZa+czfJ/ZWASK80byKRycd3xxGJFhplHdVCaYRlQ=', 1, 2);
INSERT INTO `321_user` VALUES (3, 'youyou', 'AQp2dk1ncGVWYjFDz7B0Gh+jFIsUwkV0KJs5LerYmpIzRFp31kIckwJaucc=', 1, 3);
INSERT INTO `321_user` VALUES (4, 'demo', 'AQprZUk2eHR4OHpo2WfbJU0+nqB1qyEjLBaziIxTV6sCQ55aUZaLHzpryI8=', 0, 4);
INSERT INTO `321_user` VALUES (5, 'root', 'AQpnUVVWZUxRdWx3KkvVlYRGRiJ+Wgd+6/E6skciRFuUB1614gCFNAM1YEo=', 0, 5);
INSERT INTO `321_user` VALUES (6, 'saneryao', 'AQpVNUE3N1loRTVQqtmyV3hLnrmVyLBmSQZweFn4OpxL/SeEQ3mmHU7ar+s=', 1, 6);
INSERT INTO `321_user` VALUES (29, 'haha', 'AQpwbUUyNEpMeE1UZ/RahaTHijR2wTofGmz8YV4Amei7MSM/aay69ZUF6M0=', 1, 29);
INSERT INTO `321_user` VALUES (39, 'abcd', 'AQoydnprMXE0WkVkR7BqhYHRaBvJvRqmNZ7wLEKbXdfGTEHg+Yy/wg70N/4=', 0, 39);
INSERT INTO `321_user` VALUES (40, 'aaaa', 'AQpFOHBDZ2pkOE1ybr92rbEotAmFPW2Za/HTnggHbzSVMxyzV1I4OOr75nE=', 0, 40);
INSERT INTO `321_user` VALUES (44, 'test1', 'AQpJUk9uZVVuSWZzpc2Nrmz2s+RKsq5YP0+LA2CpgdlPplW+fIZXWuibFwI=', 0, 44);

-- ----------------------------
-- Table structure for 321_user_role
-- ----------------------------
DROP TABLE IF EXISTS `321_user_role`;
CREATE TABLE `321_user_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT 0,
  `role_id` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Fixed;

SET FOREIGN_KEY_CHECKS = 1;
