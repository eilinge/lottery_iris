drop database if exists lottery;
create database lottery character set utf8;
use lottery

drop table if exists blackip;
drop table if exists code;
drop table if exists gift;
drop table if exists result;
drop table if exists user;
drop table if exists userday;

create table gift
(
    id              int not null primary key auto_increment,
    title           varchar(255),   -- 奖品名称
    prize_num       int,   -- 奖品数量, 0 无限量, >0 限量, <0 无奖品
    left_num        int,    -- 剩余奖品数量
    prize_code      varchar(50),  -- 0-9999 表示100%, 0-0表示万分之一的中奖概率
    prize_time      int,  -- 发奖周期, D天
    img             varchar(255),   --  奖品图片
    displayorder    int,    -- 位置序号, 小的排在前面
    gtype           int,   -- 奖品类型, 0 虚拟币, 1 虚拟券, 2 实物-小奖, 3 实物-大奖
    gdata           varchar(255),  -- 扩展数据, 如: 虚拟币数量
    time_begin      int,  -- 开始时间
    time_end        int,    -- 结束时间
    prize_data      mediumtext,   -- 发奖计划, [[时间1, 数量1],[时间2, 数量 2]]
    prize_begin     int, -- 发奖周期的开始
    prize_end       int,   -- 发奖周期的结束
    sys_status      smallint, -- 状态: 0 正常, 1 删除
    sys_created     int, -- 创建时间
    sys_updated     int, -- 修改时间
    sys_ip          varchar(50)   -- 操作人ip
);

alter table gift comment '奖品表';


create table code 
(
    id              int primary key not null auto_increment,
    gift_id         int, -- 奖品id, 关联gift
    code            varchar(255), -- 虚拟券编码
    sys_created     int, -- 创建时间
    sys_updated     int, -- 更新时间
    sys_status      smallint -- 状态: 0 正常, 1 作废, 2 已发放
);

create table result
(
    id              int primary key not null auto_increment,
    gift_id         int, -- 奖品id, 关联gift
    gift_name       varchar(255), -- 奖品名称
    gift_type       int, -- 奖品类型, 同gift.gtype
    uid             int,    -- 用户ID
    username        varchar(255), -- 用户名
    prize_code      int, -- 抽奖编号(4位随机数)
    gift_data       varchar(255), -- 获奖信息
    sys_created     int, -- 创建时间
    sys_ip          int, -- 用户抽奖的ip
    sys_status      smallint -- 状态: 0 正常, 1 删除, 2 作弊
);

create table blackip
(
    id              int primary key not null auto_increment,
    username        varchar(255), -- 用户名
    blacktime       int, -- 黑名单限制到期时间
    realname        varchar(50), -- 联系人
    mobile          varchar(50), -- 手机号
    address         varchar(255), -- 联系地址
    sys_created     int, -- 创建时间
    sys_updated     int, -- 修改时间
    sys_ip          varchar(50) -- IP地址
);

create table user
(
    id              int primary key not null auto_increment, 
    ip              varchar(50), -- ip 地址
    blacktime       int, -- 黑名单限制到期时间
    sys_created     int, -- 创建时间
    sys_updated     int -- 修改时间
);

-- 用户每日次数表
create table userday
(
    id              int primary key not null auto_increment,
    uid             int, -- 用户ID
    day             int, -- 日期, 如:20180725
    num             int, -- 次数                                                                                                                            
    sys_created     int, -- 创建时间
    sys_updated     int -- 修改时间
);

-- alter table vote comment '投票表，一个账户一个图片，只能投一票，一票代表30pxc';
-- CREATE UNIQUE INDEX vote_uindex ON lottery.vote (address,content_hash);

delete from code;
delete from gift;
delete from result;
delete from user;
delete from userday;
delete from blackip;