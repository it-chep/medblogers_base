-- +goose Up
-- +goose StatementBegin
alter table blog
    add column doctor_id bigint;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table blog
    drop column doctor_id;
-- +goose StatementEnd
