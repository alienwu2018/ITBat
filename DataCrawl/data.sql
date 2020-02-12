drop table if exists book_crawl;
create table book_crawl
(
    `Id`           bigint(12)    not null auto_increment comment '书本id',
    `BookName`     varchar(255)  comment '书名',
    `PngUrl`       varchar(255)  comment '书籍图片链接',
    `Author`       char(50)      comment '作者',
    `Press`        char(50)      comment '出版社',
    `PublishTime`  char(50)      comment '出版时间',
    `Pages`        int(20)       not null default 0 comment '页数',
    `ISBN`         char(50)      comment 'ISBN',
    `DouBanUrl`    varchar(255)  comment '豆瓣链接',
    `DouBanScore`  float(50)     not null default 0.0 comment '豆瓣评分',
    `Content`      text          comment '内容简介',
    `BuyUrl`       text  comment '购买链接',
    `YunPanUrl`    varchar(255)  comment '云盘链接',
    `BigCategory`  char(50)      comment '大类别',
    `SmallCategory` char(50)     comment '小类别',
    primary key (`Id`),
    key (`BookName`),
    key (`DouBanScore`)
)engine =innodb auto_increment =1 default charset=utf8 comment '书籍书籍表';

drop table  if exists downLink_crawl;
create table downLink_crawl
(
    `Id`          bigint(12)      not null auto_increment comment '自增id',
    `bid`         bigint(12)      not null comment '书本id',
    `purl`        varchar(255)    not null comment '云盘链接',
    `pwd`         char(20)        not null comment '提取码',
    primary key (`Id`)
    #constraint b_2_d_fk foreign key (bid) references book(id)
)engine =innodb auto_increment =1 default charset=utf8 comment '书籍下载信息';