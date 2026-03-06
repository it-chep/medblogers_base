-- +goose Up
-- +goose StatementBegin
-- +goose NO TRANSACTION
create index concurrently idx_operation_doctor_id_occurred on mbc_operation (doctor_id, occurred_at desc);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index idx_operation_doctor_id_occurred;
-- +goose StatementEnd
