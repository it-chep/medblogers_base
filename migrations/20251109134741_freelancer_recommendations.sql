-- +goose Up
-- +goose StatementBegin
create table if not exists freelancer_recommendation
(
    id            bigserial,
    freelancer_id bigint, -- фрилансер
    doctor_id     bigint, -- доктор, который может порекомендовать фрилансера

    unique (freelancer_id, doctor_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists freelancer_recommendation;
-- +goose StatementEnd
