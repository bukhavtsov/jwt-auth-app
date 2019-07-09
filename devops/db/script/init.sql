create table developers
(
    id            serial      not null
        constraint developers_pk
            primary key,
    name          varchar(20) not null,
    age           integer     not null,
    primary_skill varchar(20) not null
);

alter table developers
    owner to postgres;

create table users
(
    id            serial                   not null
        constraint users_pk
            primary key,
    login         varchar(20)              not null,
    password      varchar(20)              not null,
    email         varchar(30)              not null,
    role          varchar(10) default USER not null,
    refresh_token varchar(10000)
);

alter table users
    owner to postgres;

create unique index users_login_uindex
    on users (login);

