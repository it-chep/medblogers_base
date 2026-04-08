-- +goose Up
create table if not exists cookie_users
(
    cookie_id        uuid primary key,
    last_activity_at timestamp,
    domain_name      text
);

-- +goose Down
drop table if exists cookie_users;
