-- +goose Up
-- +goose StatementBegin
-- +goose NO TRANSACTION
create index concurrently if not exists idx_fr_recommendations_freelancer_id on freelancer_recommendation (freelancer_id);
create index concurrently if not exists idx_fr_recommendations_doctor_id on freelancer_recommendation (doctor_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_fr_recommendations_freelancer_id;
drop index if exists idx_fr_recommendations_doctor_id;
-- +goose StatementEnd
