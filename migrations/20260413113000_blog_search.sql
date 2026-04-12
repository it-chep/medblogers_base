-- +goose Up
-- +goose StatementBegin
alter table blog
    add column if not exists search_text text,
    add column if not exists search_vector tsvector;

create index if not exists idx_blog_search_vector
    on blog
    using gin (search_vector);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_blog_search_vector;

alter table blog
    drop column if exists search_vector,
    drop column if exists search_text;
-- +goose StatementEnd
