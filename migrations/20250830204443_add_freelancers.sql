-- +goose Up
-- +goose StatementBegin

-- Справочник специальность фрилансера
create table if not exists freelancers_speciality
(
    id   bigserial primary key,
    name varchar(255) not null -- название специальности фрилансера
);

-- Справочник город фрилансера
create table if not exists freelancers_city
(
    id   bigserial primary key,
    name varchar(255) not null -- название города фрилансера
);

-- Справочник Социальные сети
create table if not exists social_networks
(
    id   bigserial primary key,
    name varchar(30) not null -- название соцсети
);

-- Тип размещения
create table if not exists freelancers_cooperation_type
(
    id   serial primary key,
    name text -- Название типа
);

-- Фрилансер
create table if not exists freelancer
(
    id                     bigserial primary key, -- id фрилансера в системе
    email                  varchar(255) not null, -- email фрилансера
    slug                   text         not null,
    name                   varchar(255) not null,
    is_worked_with_doctors bool default false,    -- есть ли опыт работы с врачами
    is_active              bool default false,    -- признак активности
    tg_username            varchar(255),          -- ссылка на личный тг для связи
    portfolio_link         varchar(255),          -- ссылка на портфолио
    speciality_id          bigint,                -- id основной специальности фрилансера
    city_id                bigint,                -- id основного города фрилансера
    price_category         int,                   -- ценовая категория фрилансера
    s3_image               text,                  -- фотография фрилансера
    avatar                 varchar(100),          -- фотография для обратной совместимости
    has_command            bool,                  -- флаг наличия команды
    start_working_date     timestamp,             -- время начала работа на фрилансе для подсчета опыта
    cooperation_type_id    int,                   -- тип размещения

    unique (name, email),

    -- Foreign Keys
    constraint fk_freelancer_speciality
        foreign key (speciality_id)
            references freelancers_speciality (id)
            on delete set null,

    constraint fk_freelancer_city
        foreign key (city_id)
            references freelancers_city (id)
            on delete set null,

    constraint fk_freelancer_cooperation_type
        foreign key (cooperation_type_id)
            references freelancers_cooperation_type (id)
            on delete set null
);

-- M2M специальность фрилансера
create table if not exists freelancer_speciality_m2m
(
    id            bigserial primary key,
    speciality_id bigint not null,
    freelancer_id bigint not null,

    unique (speciality_id, freelancer_id),

    -- Foreign Keys
    constraint fk_speciality_m2m_speciality
        foreign key (speciality_id)
            references freelancers_speciality (id)
            on delete cascade,

    constraint fk_speciality_m2m_freelancer
        foreign key (freelancer_id)
            references freelancer (id)
            on delete cascade
);

-- M2M город фрилансера
create table if not exists freelancer_city_m2m
(
    id            bigserial primary key,
    city_id       bigint not null,
    freelancer_id bigint not null,

    unique (city_id, freelancer_id),

    -- Foreign Keys
    constraint fk_city_m2m_city
        foreign key (city_id)
            references freelancers_city (id)
            on delete cascade,

    constraint fk_city_m2m_freelancer
        foreign key (freelancer_id)
            references freelancer (id)
            on delete cascade
);

-- Социальные сети фрилансера
create table if not exists freelancer_social_networks_m2m
(
    id                bigserial primary key,
    social_network_id bigint not null,
    freelancer_id     bigint not null,

    unique (social_network_id, freelancer_id),

    -- Foreign Keys
    constraint fk_social_m2m_network
        foreign key (social_network_id)
            references social_networks (id)
            on delete cascade,

    constraint fk_social_m2m_freelancer
        foreign key (freelancer_id)
            references freelancer (id)
            on delete cascade
);

-- Прайс-лист фрилансера
create table if not exists freelancers_price_list
(
    id            bigserial primary key,
    freelancer_id bigint       not null, -- id фрилансера в системе
    name          varchar(255) not null, -- название услуги
    price         integer      not null, -- стоимость услуги. Если 0, то по договоренности

    unique (freelancer_id, name, price),

    -- Foreign Key
    constraint fk_price_list_freelancer
        foreign key (freelancer_id)
            references freelancer (id)
            on delete cascade
);

-- Создадим индексы для улучшения производительности
create index if not exists idx_freelancer_speciality on freelancer (speciality_id);
create index if not exists idx_freelancer_city on freelancer (city_id);
create index if not exists idx_freelancer_active on freelancer (is_active);
create index if not exists idx_freelancer_worked on freelancer (is_worked_with_doctors);

create index if not exists idx_speciality_m2m_speciality on freelancer_speciality_m2m (speciality_id);
create index if not exists idx_speciality_m2m_freelancer on freelancer_speciality_m2m (freelancer_id);

create index if not exists idx_city_m2m_city on freelancer_city_m2m (city_id);
create index if not exists idx_city_m2m_freelancer on freelancer_city_m2m (freelancer_id);

create index if not exists idx_social_m2m_network on freelancer_social_networks_m2m (social_network_id);
create index if not exists idx_social_m2m_freelancer on freelancer_social_networks_m2m (freelancer_id);

create index if not exists idx_price_list_freelancer on freelancers_price_list (freelancer_id);
create index if not exists idx_freelancer_cooperation_type on freelancer (cooperation_type_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- В обратном порядке удаляем таблицы (сначала дочерние, потом родительские)
drop table if exists freelancers_price_list;
drop table if exists freelancer_social_networks_m2m;
drop table if exists freelancer_city_m2m;
drop table if exists freelancer_speciality_m2m;
drop table if exists freelancers_cooperation_type;
drop table if exists freelancer;
drop table if exists social_networks;
drop table if exists freelancers_city;
drop table if exists freelancers_speciality;
-- +goose StatementEnd