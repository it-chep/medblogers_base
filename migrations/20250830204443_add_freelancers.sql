-- +goose Up
-- +goose StatementBegin

-- Справочник специальность фрилансера
create table if not exists freelancers_speciality
(
    id   bigserial,
    name varchar(255) -- название специальности фрилансера
);

-- M2M специальность фрилансера
create table if not exists freelancer_speciality_m2m
(
    id            bigserial,
    speciality_id bigint,
    freelancer_id bigint,

    unique (speciality_id, freelancer_id)
);

-- Справочник город фрилансера
create table if not exists freelancers_city
(
    id   bigserial,
    name varchar(255) -- название города фрилансера
);

-- M2M город фрилансера
create table if not exists freelancer_city_m2m
(
    id            bigserial,
    city_id       bigint,
    freelancer_id bigint,

    unique (city_id, freelancer_id)
);

-- Справочник Социальные сети
create table if not exists social_networks
(
    id   bigserial,
    name varchar(30) -- название соцсети
);

-- Социальные сети фрилансера
create table if not exists freelancer_social_networks_m2m
(
    id                bigserial,
    social_network_id bigint,
    freelancer_id     bigint,

    unique (social_network_id, freelancer_id)
);

-- Прайс-лист фрилансера
create table if not exists freelancers_price_list
(
    id            bigserial,
    freelancer_id bigint,       -- id фрилансера в системе
    name          varchar(255), -- название услуги
    price         integer,      -- стоимость услуги. Если 0, то по договоренности

    unique (freelancer_id, name, price)
);

-- Фрилансер
create table if not exists freelancer
(
    id                     bigserial,    -- id фрилансера в системе
    email                  varchar(255), -- email фрилансера
    slug                   text         not null,
    name                   varchar(255) not null,
    is_worked_with_doctors bool,         -- есть ли опыт работы с врачами
    is_active              bool,         -- признак активности
    tg_username            varchar(255), -- ссылка на личный тг для связи
    portfolio_link         varchar(255), -- ссылка на портфолио
    speciality_id          bigint,       -- id основной специальности фрилансера
    city_id                bigint,       -- id основного города фрилансера
    price_category         int,          -- ценовая категория фрилансера, определяется на основе прайс-листа или руками
    s3_image               text,         -- фотография фрилансера

    unique (name, email)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists freelancers_speciality;
drop table if exists freelancer_speciality_m2m;
drop table if exists freelancers_city;
drop table if exists freelancer_city_m2m;
drop table if exists freelancers_price_list;
drop table if exists freelancer;
drop table if exists social_networks;
drop table if exists freelancer_social_networks_m2m;
-- +goose StatementEnd
