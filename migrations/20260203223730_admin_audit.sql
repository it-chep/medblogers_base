-- +goose Up
-- +goose StatementBegin

-- История действий пользователя в админке
create table if not exists admin_audit
(
    id          uuid primary key default gen_random_uuid(),
    created_at  timestamp        default now(), -- Дата создания
    user_id     bigint not null,                -- Юзер выполнивший действие
    description text,                           -- Описание - старый доктор
    body        jsonb,                          -- Тело события. {"name": "Андрей", "inst_url": "https://inst.com"}
    action      text,                           -- Действие. Пример - Обновление данных врача
    entity_name text,                           -- Название сущности. Пример: doctor, freelancer, getcourse_user
    entity_id   bigint                          -- ID сущности
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table admin_audit;
-- +goose StatementEnd
