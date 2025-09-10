-- +goose Up
-- +goose StatementBegin
alter table docstar_site_doctor
    drop constraint if exists unique_doctor_name_email;

alter table docstar_site_doctor
    add constraint unique_doctor_name_email UNIQUE (name, email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table docstar_site_doctor
    drop constraint if exists unique_doctor_name_email;
-- +goose StatementEnd
