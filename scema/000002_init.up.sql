create table users
(
    id serial not null primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    image text not null,
    age int not null,
    weight int not null,
    height int not null
);

create table test
(
    id serial not null primary key,
    username varchar(255) not null
);