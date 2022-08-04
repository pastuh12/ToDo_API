CREATE TABLE users (
    id bigserial not null primary key,
    enconding_password varchar not null,
    email_address varchar(50) not null unique,
    name varchar(25) not null
);