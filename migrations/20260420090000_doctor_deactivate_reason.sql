-- +goose Up
alter table docstar_site_doctor
    add column if not exists deactivate_reason text;

-- +goose Down
alter table docstar_site_doctor
    drop column if exists deactivate_reason;
