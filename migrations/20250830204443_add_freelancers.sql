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
    freelancer_id bigint
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
    freelancer_id bigint
);

-- Прайс-лист фрилансера
create table if not exists freelancers_price_list
(
    id            bigserial,
    freelancer_id bigint,       -- id фрилансера в системе
    name          varchar(255), -- название услуги
    price         integer       -- стоимость услуги. Если 0, то по договоренности
);

-- Фрилансер
create table if not exists freelancer
(
    id                     bigserial,    -- id фрилансера в системе
    email                  varchar(255), -- email фрилансера
    first_name             varchar(255),
    last_name              varchar(255),
    middle_name            varchar(255),
    is_worked_with_doctors bool,         -- есть ли опыт работы с врачами
    tg_username            varchar(255), -- ссылка на личный тг для связи
    portfolio_link         varchar(255), -- ссылка на портфолио
    where_work             int[],        -- соцсети в которых работает фрилансер
    speciality_id          bigint,       -- id основной специальности фрилансера
    city_id                bigint        -- id основного города фрилансера
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
-- +goose StatementEnd
