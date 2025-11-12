-- +goose Up
-- +goose NO TRANSACTION
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_fr_recommendations_freelancer_id ON freelancer_recommendation (freelancer_id);
-- +goose NO TRANSACTION
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_fr_recommendations_doctor_id ON freelancer_recommendation (doctor_id);

-- +goose Down
-- +goose NO TRANSACTION
DROP INDEX CONCURRENTLY IF EXISTS idx_fr_recommendations_freelancer_id;
-- +goose NO TRANSACTION
DROP INDEX CONCURRENTLY IF EXISTS idx_fr_recommendations_doctor_id;