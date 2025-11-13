-- +goose Up
-- +goose NO TRANSACTION
create index concurrently if not exists idx_doc_additional_spec_doc_id on docstar_site_doctor_additional_specialties (doctor_id);
create index concurrently if not exists idx_doc_additional_spec_spec_id on docstar_site_doctor_additional_specialties (speciallity_id);
create index concurrently if not exists idx_doc_additional_city_doc_id on docstar_site_doctor_additional_cities (doctor_id);
create index concurrently if not exists idx_doc_additional_city_city_id on docstar_site_doctor_additional_cities (city_id);

-- +goose Down
-- +goose NO TRANSACTION
drop index if exists idx_doc_additional_spec_doc_id;
drop index if exists idx_doc_additional_spec_spec_id;
drop index if exists idx_doc_additional_city_doc_id;
drop index if exists idx_doc_additional_city_city_id;
