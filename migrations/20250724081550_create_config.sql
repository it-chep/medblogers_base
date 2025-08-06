-- +goose Up
-- +goose StatementBegin
create table if not exists config (
    key text primary key,
    value json not null,
    description text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists config;
-- +goose StatementEnd
