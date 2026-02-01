-- +goose Up
-- +goose StatementBegin
-- +goose NO TRANSACTION
create index concurrently doctor_id_idx on blog (doctor_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index concurrently doctor_id_idx;
-- +goose StatementEnd
