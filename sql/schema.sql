create table if not exists users
(
    id       uuid default gen_random_uuid() primary key,
    email    varchar(255) not null unique,
    password varchar(255) not null,
    username varchar(64)  not null
)