-- phpMyAdmin SQL Dump
-- version 4.5.2
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1:3306
-- Generation Time: 2016-10-26 02:24:10
-- 服务器版本： 5.6.22-log
-- PHP Version: 5.6.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `weidaogou`
--

-- --------------------------------------------------------

--
-- 表的结构 `cron`
--

CREATE TABLE `cron` (
  `id` int(11) UNSIGNED NOT NULL,
  `title` varchar(48) NOT NULL,
  `desc` tinytext NOT NULL,
  `type` enum('h','d','m') NOT NULL,
  `run_script` varchar(128) NOT NULL,
  `run_time` varchar(240) NOT NULL,
  `status` tinyint(4) NOT NULL,
  `last_exec_time` int(11) NOT NULL,
  `next_exec_time` int(11) NOT NULL,
  `path` varchar(48) NOT NULL,
  `language` varchar(24) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `cron`
--

INSERT INTO `cron` (`id`, `title`, `desc`, `type`, `run_script`, `run_time`, `status`, `last_exec_time`, `next_exec_time`, `path`, `language`) VALUES
(1, '每日抓取榜单', '每天抓取各网站的榜单数据', 'd', 'http_get.py type1', '((hour==19 and minute==30))', 0, 1450870213, 1450956613, '/data/www/spider/', 'python'),
(2, '每日抓取榜单', '每天抓取各网站的榜单数据', 'd', 'http_get.py type2', '((hour==19 and minute==35))', 0, 1450870525, 1450956925, '/data/www/spider/', 'python'),
(3, '每日抓取榜单', '每天抓取各网站的榜单数据', 'd', 'http_get.py type3', '((hour==19 and minute==40))', 0, 1450870826, 1450957226, '/data/www/spider/', 'python'),
(4, '榜单分析程序', '每天分析榜单的数据', 'd', 'analyse_rank.php', '((hour==0 and minute==15))', 0, 1450887302, 1450973702, '/data/www/3dmin/you/script/', 'php'),
(5, '定时抓取APP数据更新', '将APP中的部分数据抓取后针对当前的游戏库进行数据更新', 'd', 'get_games_info_itunes.php', '((hour==2 and minute==50))', 0, 1450896615, 1450983015, '/data/www/3dmin/you/script/', 'php'),
(6, '定时更新静态页面', '主页和新闻列表页目前', 'h', 'run_static.php', '((minute==10))', 0, 1450937429, 1450941029, '/data/www/3dmin/you/script/', 'php');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `cron`
--
ALTER TABLE `cron`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `cron`
--
ALTER TABLE `cron`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
