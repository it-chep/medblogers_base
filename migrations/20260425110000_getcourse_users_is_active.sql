-- +goose Up
-- +goose StatementBegin
alter table getcourse_users
    add column if not exists is_active boolean not null default true;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table getcourse_users
    drop column if exists is_active;
-- +goose StatementEnd
