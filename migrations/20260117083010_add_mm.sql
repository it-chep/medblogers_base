-- +goose Up
-- +goose StatementBegin
create table if not exists mm
(
    id          bigserial primary key,
    mm_datetime timestamp,
    name        text,
    state       int,
    mm_link     text,
    created_at  timestamp default now(),
    is_active   bool
);

create table if not exists getcourse_orders
(
    id         bigserial primary key,
    order_id   bigint,
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


-- Рассылки
create table if not exists newsletter
(
    newsletter_uuid uuid primary key default gen_random_uuid(),
    created_at      timestamp        default now(),
    planned_sb_ids  bigint[],
    event_type      text
);

-- Отправленные рассылки
create table if not exists sent_newsletter
(
    id              bigserial primary key,
    newsletter_uuid uuid,
    sb_id           bigint,
    created_at      timestamp default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table getcourse_users;
drop table getcourse_orders;
drop table mm;
-- +goose StatementEnd
