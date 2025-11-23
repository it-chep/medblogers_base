-- +goose Up
-- +goose StatementBegin
-- Статьи
create table if not exists blog
(
    id                  uuid primary key default gen_random_uuid(),
    name                text,        -- название статьи
    slug                text unique, -- отображение по урлу
    is_active           bool,        -- показывать ее или нет
    body                text,        -- контент статьи в формате html
    preview_text        text,        -- Превью текст в списке статей
    society_preview     text,        -- превью в соц сетях
    additional_seo_text text         -- Дополнительный SEO текст
);

-- Фотографии в статьях
create table if not exists blog_photos
(
    id        uuid primary key,
    blog_id   uuid,
    file_type text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists blog;
drop table if exists blog_photos
-- +goose StatementEnd
