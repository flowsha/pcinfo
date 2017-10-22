/*
Navicat MySQL Data Transfer

Source Server         : liusha
Source Server Version : 50527
Source Host           : localhost:3306
Source Database       : zhwh

Target Server Type    : MYSQL
Target Server Version : 50527
File Encoding         : 65001

Date: 2017-10-22 19:41:26
*/

SET FOREIGN_KEY_CHECKS=0;
-- ----------------------------
-- Table structure for `tb_pcinfo`
-- ----------------------------
DROP TABLE IF EXISTS `tb_pcinfo`;
CREATE TABLE `tb_pcinfo` (
  `ID` varchar(40) NOT NULL,
  `V_SN_SHA` varchar(40) DEFAULT NULL,
  `V_DEPARTMENT` varchar(50) DEFAULT NULL,
  `V_USERNAME` varchar(20) DEFAULT NULL,
  `V_COMPUTER_NAME` varchar(50) DEFAULT NULL,
  `V_COMPUTER_MANUFACTURER` varchar(50) DEFAULT NULL,
  `V_COMPUTER_MODEL` varchar(50) DEFAULT NULL,
  `V_COMPUTER_ARCH` varchar(50) DEFAULT NULL,
  `V_COMPUTER_TYPE` varchar(10) DEFAULT NULL,
  `V_CPU_MODEL` varchar(50) DEFAULT NULL,
  `V_CPU_MANUFACTURER` varchar(50) DEFAULT NULL,
  `V_CPU_CORES` int(6) DEFAULT NULL,
  `V_MEMORY_CAPACITY` int(4) DEFAULT NULL,
  `V_MEMORY_NUM` smallint(6) DEFAULT NULL,
  `V_NETWORK1_MAC` varchar(17) DEFAULT NULL,
  `V_NETWORK1_IPV4` varchar(15) DEFAULT NULL,
  `V_NETWORK1_IPV6` varchar(29) DEFAULT NULL,
  `V_NETWORK2_MAC` varchar(17) DEFAULT NULL,
  `V_NETWORK2_IPV4` varchar(15) DEFAULT NULL,
  `V_NETWORK2_IPV6` varchar(29) DEFAULT NULL,
  `V_DISK_MODEL` varchar(50) DEFAULT NULL,
  `V_DISK_SN` varchar(50) DEFAULT NULL,
  `V_DISK_CAPACITY` varchar(10) DEFAULT NULL,
  `V_DISK_INTERFACE` varchar(10) DEFAULT NULL,
  `V_MONITOR_NAME` varchar(50) DEFAULT NULL,
  `V_MONITOR_TYPE` varchar(50) DEFAULT NULL,
  `D_PRODUCT_DATE` date DEFAULT NULL,
  `D_ADD_DATE` datetime DEFAULT NULL,
  `D_MODIFY_DATE` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_pcinfo
-- ----------------------------
