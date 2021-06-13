-- drop schema app;
create database if not exists app;

use app;

-- 用户信息
create table user_info(
    u_id    integer     not null auto_increment,
    u_name  varchar(15) not null unique,
    u_psw   varchar(35) not null,

    primary key (u_id)
);

-- 活动信息
create table act(
    act_id      integer     not null,
    org_id      integer     not null,
    act_name    varchar(15) not null,
    act_len     integer     not null,
    act_des     varchar(60) not null,
    act_stop    boolean     not null default false,
    act_final   datetime    null,

    primary key (act_id),
    foreign key (org_id) references user_info(u_id)
);

-- 活动对应的时间段
create table act_period(
    period_id   integer     not null auto_increment,
    act_id      integer     not null,
    org_period  datetime    not null,
    -- vote        integer     not null default 0,
    primary key (period_id),
    foreign key (act_id) references act(act_id)
);

-- 每个活动的每个时间段，与它的投票者
create table vote(
    period_id   integer not null,
    u_id        integer not null,

    primary key (period_id, u_id),
    foreign key (period_id) references act_period(period_id),
    foreign key (u_id) references user_info(u_id)
);

-- 活动对应的参与者
create table act_part(
    act_id  integer     not null,
    u_id    integer     not null,

    primary key (act_id, u_id),
    foreign key (act_id) references act(act_id),
    foreign key (u_id) references user_info(u_id)
);