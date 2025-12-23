-- +goose Up
-- +goose StatementBegin
alter table freelancer
    drop column has_command;

alter table freelancer
    drop column is_worked_with_doctors;

alter table freelancer
    add column agency_representative bool;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
