-- +goose Up
create table if not exists utm_analytics
(
    uuid         uuid primary key,
    domain_name  varchar(100), -- название сайта medblogers.ru, medblogers-base.ru
    created_at   timestamp default now(),
    company      varchar(100), -- промо-компания, необходимо для группировки в сквозную аналитику
    event        varchar(100), -- шаг в воронке, чтобы отследить сквозное действие пользователя (оставил контакт, перешел по боту дальше)

    cookie_id    uuid,         -- идентификатор сессии пользователя

    -- метки
    utm_source   varchar(100),
    utm_medium   varchar(100),
    utm_campaign varchar(100),
    utm_term     varchar(100),
    utm_content  varchar(100)
);

create index if not exists idx_utm_company on utm_analytics (company);
create index if not exists idx_utm_event on utm_analytics (event);
create index if not exists idx_utm_created_at on utm_analytics (created_at);

-- +goose Down
drop table if exists utm_analytics;
