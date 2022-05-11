- phpMyAdmin SQL Dump
-- version 4.6.6
-- https://www.phpmyadmin.net/
--
-- 主機: localhost
-- 產生時間： 2022-05-10 03:57:02
-- 伺服器版本: 5.7.17-log
-- PHP 版本： 5.6.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- 資料庫： `cash_register`
--

-- --------------------------------------------------------

--
-- 資料表結構 `member_info`
--

CREATE TABLE `member_info` (
  `id` int(50) NOT NULL,
  `name` varchar(50) NOT NULL COMMENT '會員名稱',
  `password` varchar(50) NOT NULL COMMENT '會員密碼',
  `email` varchar(50) NOT NULL COMMENT '會員郵件(不可重複)',
  `vip` int(50) NOT NULL COMMENT '會員VIP等級'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='會員資料';


-- --------------------------------------------------------

--
-- 資料表結構 `platform_coin`
--

CREATE TABLE `platform_coin` (
  `id` int(50) NOT NULL,
  `uid` int(50) NOT NULL COMMENT '會員id',
  `account` float NOT NULL COMMENT '帳戶餘額',
  `cost` float NOT NULL COMMENT '消費金額'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='會員平台幣帳戶餘額及消費金額';


-- --------------------------------------------------------

--
-- 資料表結構 `platform_point`
--

CREATE TABLE `platform_point` (
  `id` int(50) NOT NULL,
  `uid` int(50) NOT NULL COMMENT '會員id',
  `account` float NOT NULL COMMENT '帳戶點數',
  `cost` float NOT NULL COMMENT '消費點數'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='會員平台點數帳戶點數及消費點數';


-- --------------------------------------------------------

--
-- 資料表結構 `platform_point_percentage`
--

CREATE TABLE `platform_point_percentage` (
  `percentage` float NOT NULL COMMENT '平台點數折抵平台幣比率'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='扣平台點數折抵平台幣比率';

--
-- 資料表的匯出資料 `platform_point_percentage`
--


-- --------------------------------------------------------

--
-- 資料表結構 `vip_coin_discount`
--

CREATE TABLE `vip_coin_discount` (
  `vip` int(50) NOT NULL,
  `disacount` float NOT NULL COMMENT 'VIP等級平台幣優惠比率'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='VIP會員平台幣優惠比率';

--
-- 資料表的匯出資料 `vip_coin_discount`
--


-- --------------------------------------------------------

--
-- 資料表結構 `vip_point_discount`
--

CREATE TABLE `vip_point_discount` (
  `discount` float NOT NULL COMMENT 'VIP身份扣100點以上折抵比率'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='VIP身份扣100點以上折抵比率';


--
-- 已匯出資料表的索引
--

--
-- 資料表索引 `member_info`
--
ALTER TABLE `member_info`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `vip` (`vip`);

--
-- 資料表索引 `platform_coin`
--
ALTER TABLE `platform_coin`
  ADD PRIMARY KEY (`id`),
  ADD KEY `uid` (`uid`);

--
-- 資料表索引 `platform_point`
--
ALTER TABLE `platform_point`
  ADD PRIMARY KEY (`id`),
  ADD KEY `uid` (`uid`);

--
-- 資料表索引 `vip_coin_discount`
--
ALTER TABLE `vip_coin_discount`
  ADD PRIMARY KEY (`vip`);

--
-- 在匯出的資料表使用 AUTO_INCREMENT
--

--
-- 使用資料表 AUTO_INCREMENT `member_info`
--
ALTER TABLE `member_info`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT;
--
-- 使用資料表 AUTO_INCREMENT `platform_coin`
--
ALTER TABLE `platform_coin`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT;
--
-- 使用資料表 AUTO_INCREMENT `platform_point`
--
ALTER TABLE `platform_point`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT;
--
-- 已匯出資料表的限制(Constraint)
--

--
-- 資料表的 Constraints `member_info`
--
ALTER TABLE `member_info`
  ADD CONSTRAINT `member_info_ibfk_1` FOREIGN KEY (`vip`) REFERENCES `vip_coin_discount` (`vip`);

--
-- 資料表的 Constraints `platform_coin`
--
ALTER TABLE `platform_coin`
  ADD CONSTRAINT `platform_coin_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `member_info` (`id`);

--
-- 資料表的 Constraints `platform_point`
--
ALTER TABLE `platform_point`
  ADD CONSTRAINT `platform_point_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `member_info` (`id`);
