CREATE TABLE tasks (
    id bigserial not null primary key,
    title varchar(50) not null,
    description varchar(200) not null, 
    status status_type DEFAULT 'not done',
    folder_id bigint
);