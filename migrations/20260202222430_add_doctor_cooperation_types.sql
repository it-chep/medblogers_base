-- +goose Up
-- +goose StatementBegin
create table if not exists doctors_cooperation_type
(
    id   bigserial primary key,
    name text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists doctors_cooperation_type;
-- +goose StatementEnd
