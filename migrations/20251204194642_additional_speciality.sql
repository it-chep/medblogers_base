-- +goose Up
-- +goose StatementBegin

-- Дополнительные специальности
create table if not exists additional_medical_specialities
(
    id                       bigserial primary key,
    primary_speciality_id    bigint,
    additional_speciality_id bigint
);

alter table docstar_site_speciallity
    add column is_only_additional bool;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists additional_medical_specialities;

alter table docstar_site_speciallity
    drop column is_only_additional;
-- +goose StatementEnd
