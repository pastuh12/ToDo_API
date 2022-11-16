CREATE TABLE sessions(
    id bigserial not null primary key,
    userID bigint not null unique,
    refreshToken varchar(200) not null unique,
    expiresAt int not null,
    CONSTRAINT fk_user
      FOREIGN KEY(userID) 
	  REFERENCES users(id)
      ON DELETE CASCADE
);