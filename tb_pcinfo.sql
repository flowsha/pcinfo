/*
Navicat MySQL Data Transfer

Source Server         : xtcdma
Source Server Version : 50075
Source Host           : 134.162.72.226:3306
Source Database       : zhwh

Target Server Type    : MYSQL
Target Server Version : 50075
File Encoding         : 65001

Date: 2017-10-23 15:55:43
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `tb_pcinfo`
-- ----------------------------
DROP TABLE IF EXISTS `tb_pcinfo`;
CREATE TABLE `tb_pcinfo` (
  `ID` varchar(40) NOT NULL,
  `V_SN_SHA` varchar(40) default NULL,
  `V_DEPARTMENT` varchar(50) default NULL,
  `V_USERNAME` varchar(20) default NULL,
  `V_COMPUTER_NAME` varchar(50) default NULL,
  `V_COMPUTER_MANUFACTURER` varchar(50) default NULL,
  `V_COMPUTER_MODEL` varchar(50) default NULL,
  `V_COMPUTER_ARCH` varchar(50) default NULL,
  `V_COMPUTER_TYPE` varchar(10) default NULL,
  `V_CPU_MODEL` varchar(50) default NULL,
  `V_CPU_MANUFACTURER` varchar(50) default NULL,
  `V_CPU_CORES` int(6) default NULL,
  `V_MEMORY_CAPACITY` int(4) default NULL,
  `V_MEMORY_NUM` smallint(6) default NULL,
  `V_NETWORK1_MAC` varchar(17) default NULL,
  `V_NETWORK1_IPV4` varchar(15) default NULL,
  `V_NETWORK1_IPV6` varchar(29) default NULL,
  `V_NETWORK2_MAC` varchar(17) default NULL,
  `V_NETWORK2_IPV4` varchar(15) default NULL,
  `V_NETWORK2_IPV6` varchar(29) default NULL,
  `V_DISK_MODEL` varchar(50) default NULL,
  `V_DISK_SN` varchar(50) default NULL,
  `V_DISK_CAPACITY` varchar(10) default NULL,
  `V_DISK_INTERFACE` varchar(10) default NULL,
  `V_MONITOR_NAME` varchar(50) default NULL,
  `V_MONITOR_TYPE` varchar(50) default NULL,
  `V_PRINTER_NAME` varchar(50) default NULL,
  `V_BIOS_NAME` varchar(50) default NULL,
  `V_BIOS_MANUFACTURER` varchar(50) default NULL,
  `V_BIOS_SN` varchar(50) default NULL,
  `D_BIOS_RELEASE_DATE` date default NULL,
  `D_INSERT_TIME` datetime default NULL,
  `D_UPDATE_TIME` datetime default NULL,
  PRIMARY KEY  (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tb_pcinfo
-- ----------------------------
