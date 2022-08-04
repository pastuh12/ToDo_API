CREATE TABLE sessions(
    id bigserial not null primary key,
    userID bigint not null unique,
    refreshToken varchar(200) not null unique,
    expiresAt int not null
);