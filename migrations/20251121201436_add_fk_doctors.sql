-- +goose Up
-- +goose StatementBegin
alter table docstar_site_doctor
    add column is_kf_doctor bool;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table docstar_site_doctor
    drop column is_kf_doctor;
-- +goose StatementEnd
