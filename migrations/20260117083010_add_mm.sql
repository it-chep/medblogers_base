-- +goose Up
-- +goose StatementBegin
create table if not exists mm
(
    id          bigserial primary key,
    mm_datetime timestamp,
    name        text,
    state       int,
    mm_link     text,
    created_at  timestamp default now()
);

create table if not exists getcourse_orders
(
    id         bigserial primary key,
    order_id   text,
    gk_id      bigint,
    days_count bigint,
    name       text,
    created_at timestamp default now()
);

create table if not exists getcourse_users
(
    id         bigserial primary key,
    sb_id      bigint,
    gk_id      bigint,
    name       bigint,
    end_date   timestamp,
    days_count bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table getcourse_users;
drop table getcourse_orders;
drop table mm;
-- +goose StatementEnd
