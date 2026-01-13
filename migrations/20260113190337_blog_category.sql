-- +goose Up
-- +goose StatementBegin
create table if not exists blog_category
(
    id   bigserial primary key,
    name text
);

create table if not exists m2m_blog_category
(
    id          bigserial primary key,
    blog_id     uuid not null,
    category_id bigint not null,

    constraint uq_blog_category unique (blog_id, category_id)
);

-- create table if not exists m2m_blog_category
-- (
--     blog_id     uuid,   -- id блога
--     category_id bigint, -- id категории
--
--     primary key (blog_id, category_id)
-- );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists blog_category;
drop table if exists m2m_blog_category
-- +goose StatementEnd
