create table if not exists `order`
(
    id int auto_increment
        primary key,
    user_id int not null,
    symbol varchar(16) null,
    order_id int null,
    client_order_id int null,
    transact_time timestamp not null,
    price varchar(18) null
);

create table if not exists user
(
    id int auto_increment
        primary key,
    api_key varchar(64) not null,
    secret_key varchar(64) not null
);

