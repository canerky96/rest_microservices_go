create table users
(
    id           bigint auto_increment
        primary key,
    first_name   varchar(255) null,
    last_name    varchar(255) null,
    email        varchar(255) not null,
    date_created varchar(255) not null,
    status       varchar(45)  not null,
    password     varchar(255) not null,

        constraint users_email_uindex
            unique (email)
);
