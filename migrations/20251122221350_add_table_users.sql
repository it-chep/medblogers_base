-- +goose Up
-- +goose StatementBegin

-- Роли
create table if not exists roles
(
    id          bigserial,
    name        varchar(50) not null, -- АНГЛ Название роли
    description text                  -- Описание роли
);

-- Права доступа
create table if not exists permissions
(
    id          bigserial,
    name        varchar(50) not null, -- Название права (уникальное)
    url         text        not null, -- Урл в системе
    description text                  -- Описание права (необязательно)
);

-- Права для ролей
create table if not exists roles_permissions
(
    id            bigserial,
    role_id       bigint, -- Роль
    permission_id bigint  -- Правило
);


create table if not exists go_users
(
    id        bigserial primary key,
    email     varchar(255) not null unique, -- email
    password  varchar(255),                 -- пароль
    full_name VARCHAR(100),                 -- фио
    is_active bool,                         -- активный ли пользователь
    role_id   bigint                        -- ID роли (админ, суперадмин)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists roles;
drop table if exists permissions;
drop table if exists roles_permissions;
drop table if exists go_users;
-- +goose StatementEnd
