-- +goose Up
-- +goose StatementBegin
alter table docstar_site_speciallity
    add column primary_speciality_id bigint;

alter table docstar_site_speciallity
    add column is_only_additional bool;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table docstar_site_speciallity
    drop column primary_speciality_id;

alter table docstar_site_speciallity
    drop column is_only_additional;
-- +goose StatementEnd
