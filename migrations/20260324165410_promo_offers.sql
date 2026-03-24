-- +goose Up
-- +goose StatementBegin

-- Справочник типа сотрудничества для офферов
create table if not exists promo_offer_cooperation_type
(
    id   bigserial primary key,
    name varchar(100) not null unique -- название типа сотрудничества
);

-- Справочник тематики офферов и брендов
create table if not exists promo_offer_topic
(
    id   bigserial primary key,
    name varchar(100) not null unique -- название тематики
);

-- Справочник форматов контента/размещения
create table if not exists promo_offer_content_format
(
    id   bigserial primary key,
    name varchar(100) not null unique -- название формата контента
);

-- Бренды
create table if not exists brand
(
    id          bigserial primary key,
    photo       text,                    -- фото бренда/компании
    title       text,                    -- заголовок
    slug        text not null unique,    -- slug бренда
    topic_id    bigint,                  -- тема компании
    website     text,                    -- сайт компании
    description text,                    -- описание компании
    is_active   bool      default false, -- признак активности
    created_at  timestamp default now()  -- дата создания
);

-- Социальные сети бренда/компании
create table if not exists brand_social_networks
(
    id                bigserial primary key,
    brand_id          bigint not null,
    social_network_id bigint not null,
    link              text   not null, -- ссылка на соцсеть компании

    unique (brand_id, social_network_id, link)
);

-- Офферы
create table if not exists promo_offer
(
    id                     uuid primary key default gen_random_uuid(),
    cooperation_type_id    bigint,                  -- тип сотрудничества
    topic_id               bigint,                  -- тема оффера
    title                  text not null,           -- заголовок
    description            text,                    -- описание, что надо сделать и что прорекламировать
    price                  bigint,                  -- цена
    content_format_id      bigint,                  -- вид размещения / формат контента
    brand_id               bigint,                  -- бренд
    publication_date       timestamp,               -- дата публикации
    ad_marking_responsible text,                    -- кто маркирует рекламу
    responses_capacity     integer   default 0,     -- объем откликов
    is_active              bool      default false, -- признак активности
    created_at             timestamp default now()  -- дата создания
);

-- Социальные сети оффера
create table if not exists promo_offer_social_networks_m2m
(
    id                bigserial primary key,
    promo_offer_id    uuid not null,
    social_network_id bigint not null,

    unique (promo_offer_id, social_network_id)
);

insert into promo_offer_cooperation_type (name)
values ('Бартер'),
       ('Амбассадор'),
       ('Разовое')
on conflict (name) do nothing;

insert into promo_offer_topic (name)
values ('Фармакология'),
       ('Блогер'),
       ('Мед. одежда'),
       ('Клиника'),
       ('Медоборудование')
on conflict (name) do nothing;

insert into promo_offer_content_format (name)
values ('Рилс'),
       ('Карусель'),
       ('Стори'),
       ('Пост'),
       ('Любой контент')
on conflict (name) do nothing;

create index if not exists idx_brand_topic on brand (topic_id);

create index if not exists idx_brand_social_networks_brand on brand_social_networks (brand_id);
create index if not exists idx_brand_social_networks_network on brand_social_networks (social_network_id);

create index if not exists idx_promo_offer_cooperation_type on promo_offer (cooperation_type_id);
create index if not exists idx_promo_offer_topic on promo_offer (topic_id);
create index if not exists idx_promo_offer_content_format on promo_offer (content_format_id);
create index if not exists idx_promo_offer_brand on promo_offer (brand_id);
create index if not exists idx_promo_offer_created_at on promo_offer (created_at desc);
create index if not exists idx_promo_offer_publication_date on promo_offer (publication_date desc);

create index if not exists idx_promo_offer_social_offer on promo_offer_social_networks_m2m (promo_offer_id);
create index if not exists idx_promo_offer_social_network on promo_offer_social_networks_m2m (social_network_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists promo_offer_social_networks_m2m;
drop table if exists brand_social_networks;
drop table if exists promo_offer;
drop table if exists brand;
drop table if exists promo_offer_content_format;
drop table if exists promo_offer_topic;
drop table if exists promo_offer_cooperation_type;
-- +goose StatementEnd
