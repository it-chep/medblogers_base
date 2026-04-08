-- +goose Up
-- +goose NO TRANSACTION
-- Создание таблицы просмотров статьи
create table if not exists blogs_views
(
    id         uuid primary key,
    blog_uuid  uuid,
    cookie_id  uuid,
    created_at timestamp default now()
);

-- Индекс на статьи
create index concurrently if not exists idx_blogs_views_uuid on blogs_views (blog_uuid);

-- +goose Down
-- +goose NO TRANSACTION
drop table if exists blogs_views;
