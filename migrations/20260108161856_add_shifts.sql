-- +goose Up
-- +goose StatementBegin

alter table docstar_site_doctor
    add column tg_id bigint;

create table if not exists patients
(
    id            bigserial primary key,
    phone         text,
    tg_id         bigint,
    name          text,
    communication text -- как связаться какая-то ссылка
);

create table if not exists doctor_slots_capacity
(
    slot_id    uuid primary key,
    doctor_id  bigint,
    slot_start timestamp,
    slot_end   timestamp,
    timezone   text
);

create table if not exists doctor_booked_slots
(
    shift_id   uuid primary key,
    doctor_id  bigint,
    patient_id bigint,
    slot_id    uuid,
    slot_start timestamp,
    slot_end   timestamp,
    created_at timestamp default now(),
    state      int
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists patients;
drop table if exists doctor_slots_capacity;
drop table if exists doctor_booked_slots;
-- +goose StatementEnd
