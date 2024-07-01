CREATE TABLE IF NOT EXISTS tasks(
    id varchar(255)  primary key,
    task varchar(255) not null,
    start_at varchar(255) not null
);