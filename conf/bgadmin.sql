-- ----------------------------
-- Source Server Type    : MySQL
-- Source Server Version : 50721

-- Target Server Type    : MySQL
-- Target Server Version : 50721
-- File Encoding         : 65001

-- Date: 12/10/2018 10:59:37
-- ----------------------------

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
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 501 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_code
-- ----------------------------
INSERT INTO `321_code` (`id`, `user_id`, `code`) VALUES
    (1, 1, '0lEHGWmsysP92EMim1M6faN7QBl4t2KaRrp6hUFtk4fGOUffEviAxrkfinMVLPHi'),
    (2, 2, 'YUbQyK5byDpMQ2pJq6SXwGeEfB5aBk9DV6gYESwSWgp8DHupzrHUBkiylpR74WQ5'),
    (3, 3, 'dBYNDmhnvcEtdC3md9NTwXfhvQzM0uR4sMzkionNcNSuFNBRJDarZ1IeKOqRBBfi'),
    (4, 4, 'W3xaWYzkzdolZwvWhAHf2cKXTIKNodVXFPdkD72V9nZYAh7jr6hwdHkwm1KoqlLc');

-- ----------------------------
-- Table structure for 321_user
-- ----------------------------
DROP TABLE IF EXISTS `321_user`;
CREATE TABLE `321_user`  (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `pwd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `state` int(11) NOT NULL DEFAULT 0,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `profile_id` bigint(20) NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `profile_id`(`profile_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 501 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_user
-- ----------------------------
INSERT INTO `321_user` (`id`, `name`, `pwd`, `state`, `profile_id`) VALUES
    (1, 'saneryao', 'AQpVNUE3N1loRTVQqtmyV3hLnrmVyLBmSQZweFn4OpxL/SeEQ3mmHU7ar+s=', 1, 1),
    (2, 'admin', 'AQpyZkJkNTZ0aTJTiMp2eLHYq5Ex+roSKU2p58M2sKp2dt5fmUBi6x8+MJY=', 0, 2),
    (3, 'test', 'AQpoUGprU0J2UU9WCcWZa+czfJ/ZWASK80byKRycd3xxGJFhplHdVCaYRlQ=', 0, 3),
    (4, 'youyou', 'AQp2dk1ncGVWYjFDz7B0Gh+jFIsUwkV0KJs5LerYmpIzRFp31kIckwJaucc=', 0, 4),
    (5, 'demo', 'AQprZUk2eHR4OHpo2WfbJU0+nqB1qyEjLBaziIxTV6sCQ55aUZaLHzpryI8=', 0, 5),
    (6, 'root', 'AQpnUVVWZUxRdWx3KkvVlYRGRiJ+Wgd+6/E6skciRFuUB1614gCFNAM1YEo=', 0, 6),
    (7, 'haha', 'AQpwbUUyNEpMeE1UZ/RahaTHijR2wTofGmz8YV4Amei7MSM/aay69ZUF6M0=', 0, 7),
    (8, 'abcd', 'AQoydnprMXE0WkVkR7BqhYHRaBvJvRqmNZ7wLEKbXdfGTEHg+Yy/wg70N/4=', 0, 8),
    (9, 'aaaa', 'AQpFOHBDZ2pkOE1ybr92rbEotAmFPW2Za/HTnggHbzSVMxyzV1I4OOr75nE=', 0, 9),
    (10, 'test1', 'AQpJUk9uZVVuSWZzpc2Nrmz2s+RKsq5YP0+LA2CpgdlPplW+fIZXWuibFwI=', 0, 10);

-- ----------------------------
-- Table structure for 321_profile
-- ----------------------------
DROP TABLE IF EXISTS `321_profile`;
CREATE TABLE `321_profile`  (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `sex` int(11) NULL DEFAULT NULL,
    `birth` date NULL DEFAULT NULL,
    `nick` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
    `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
    `mobile` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
    `tel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
    `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 501 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_profile
-- ----------------------------
INSERT INTO `321_profile` (`id`, `sex`, `birth`, `nick`, `email`, `mobile`, `tel`, `address`) VALUES
    (1, 1, NULL, 'saneryao', 'saneryao@saneryao.com', '', '', 'Xiangtan'),
    (2, 1, NULL, 'admin*', 'admin@saneryao.com', '1301106553', NULL, 'Beijing'),
    (3, 0, NULL, 'test', 'test@saneryao.com', '', '', 'Shaihai'),
    (4, 1, NULL, 'youyou', 'youyou@saneryao.com', '', '', 'Shenzhen'),
    (5, 0, NULL, 'demo', 'demo@saneryao.com', '', '', 'Guangzhou'),
    (6, 1, NULL, 'root', '498807782@qq.com', '', '', 'Hangzhou'),
    (7, 0, '2018-10-04', 'haha', 'haha@saneryao.com', '', '', 'Changsha'),
    (8, 1, '2018-10-03', '...', 'abcd@saneryao.com', '', '', ''),
    (9, 1, '2018-09-26', 'aaaa...', 'aaaa@saneryao.com', '', '', ''),
    (10, 0, NULL, '', '583233681@qq.com', '', '', '');

-- ----------------------------
-- Table structure for 321_role
-- ----------------------------
DROP TABLE IF EXISTS `321_role`;
CREATE TABLE `321_role`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `state` int(11) NOT NULL DEFAULT 0,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 501 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_role
-- ----------------------------
INSERT INTO `321_role` (`id`, `name`, `desc`, `state`) VALUES
    (1, 'admin', 'admin', 1),
    (2, 'test', 'test...', 0),
    (3, 'abcd', '...', 1),
    (4, 'aaaa', '', 0);


-- ----------------------------
-- Table structure for 321_menu
-- ----------------------------
DROP TABLE IF EXISTS `321_menu`;
CREATE TABLE `321_menu`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `parent_id` bigint(20) NOT NULL DEFAULT 0,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `state` int(11) NOT NULL DEFAULT 0,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 501 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_menu
-- ----------------------------
INSERT INTO `321_menu` (`id`, `parent_id`, `name`, `icon`, `url`, `state`) VALUES
    (1, 0, '系统首页', 'fa-home', 'page/index', 1),
    (2, 0, '用户和权限', 'fa-users', '', 1),
    (3, 2, '用户', '', 'page/user', 1),
    (4, 2, '角色', '', 'page/role', 1),
    (5, 2, '菜单', '', 'page/menu', 1),
    (6, 2, 'API', '', 'page/link', 1),
    (7, 0, '更多示例', 'fa-info-circle', '', 1),
    (8, 7, '页面', '', 'page/demo', 1),
    (9, 7, '排版', '', 'page/grid', 1),
    (10, 7, '按钮', '', 'page/button', 1),
    (11, 7, '图标', '', 'page/icon', 1);

-- ----------------------------
-- Table structure for 321_link
-- ----------------------------
DROP TABLE IF EXISTS `321_link`;
CREATE TABLE `321_link`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `power` int(11) NOT NULL DEFAULT 0,
    `state` int(11) NOT NULL DEFAULT 0,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 501 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 321_link
-- ----------------------------
INSERT INTO `321_link` (`id`, `name`, `url`, `power`, `state`) VALUES
    (1, '统计', '/api/v1/statistics', 1, 1),
    (2, '用户', '/api/v1/user', 31, 1),
    (3, '角色', '/api/v1/role', 31, 1),
    (4, '菜单', '/api/v1/menu', 31, 1),
    (5, 'API', '/api/v1/link', 31, 1);

-- ----------------------------
-- Table structure for 321_role_power
-- ----------------------------
DROP TABLE IF EXISTS `321_role_power`;
CREATE TABLE `321_role_power`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `role_id` int(20) NOT NULL DEFAULT 0,
    `menu_id` int(20) NOT NULL DEFAULT 0,
    `link_id` int(20) NOT NULL DEFAULT 0,
    `link_power` int(11) NOT NULL DEFAULT 0,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 501 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Fixed;

-- ----------------------------
-- Records of 321_role_power
-- ----------------------------
INSERT INTO `321_role_power` (`id`, `role_id`, `menu_id`, `link_id`, `link_power`) VALUES
    (1, 1, 1, 0, 0),
    (2, 1, 2, 0, 0),
    (3, 1, 3, 0, 0),
    (4, 1, 4, 0, 0),
    (5, 1, 5, 0, 0),
    (6, 1, 6, 0, 0),
    (7, 1, 7, 0, 0),
    (8, 1, 8, 0, 0),
    (9, 1, 9, 0, 0),
    (10, 1, 10, 0, 0),
    (11, 1, 11, 0, 0),
    (12, 1, 0, 1, 1),
    (13, 1, 0, 2, 31),
    (14, 1, 0, 3, 31),
    (15, 1, 0, 4, 31),
    (16, 1, 0, 5, 31);

-- ----------------------------
-- Table structure for 321_user_role
-- ----------------------------
DROP TABLE IF EXISTS `321_user_role`;
CREATE TABLE `321_user_role`  (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL DEFAULT 0,
    `role_id` int(11) NOT NULL DEFAULT 0,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 501 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Fixed;

-- ----------------------------
-- Records of 321_user_role
-- ----------------------------
INSERT INTO `321_user_role` (`id`, `user_id`, `role_id`) VALUES
    (1, 1, 1);

SET FOREIGN_KEY_CHECKS = 1;
