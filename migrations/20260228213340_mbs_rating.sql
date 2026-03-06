-- +goose Up
-- +goose StatementBegin
-- Таблица начислений mbs баллов
create table mbc_operation
(
    id            uuid        not null,               -- идентификатор начисления
    doctor_id     bigint      not null,               -- ид врача, которому начислили баллы
    mbc_count     int         not null,               -- кол-во баллов
    accrued_by_id bigint      not null,               -- кто начислили (система, ручное - логин)
    occurred_at   timestamptz not null default now(), -- время начисления
    unique (id, occurred_at)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists mbc_operation
-- +goose StatementEnd
