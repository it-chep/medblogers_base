-- +goose Up
create table if not exists breadcrumbs
(
    id        bigserial primary key,
    name      text not null,
    url       text not null unique,
    parent_id bigint
);

create index if not exists idx_breadcrumbs_parent_id on breadcrumbs (parent_id);

-- +goose Down
drop table if exists breadcrumbs;
