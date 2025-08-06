-- +goose Up
-- +goose StatementBegin
create table if not exists docstar_site_city
(
    id   bigserial primary key,
    name varchar(100) not null
);

create table if not exists docstar_site_speciallity
(
    id   bigserial primary key,
    name varchar(100) not null
);

create table if not exists docstar_site_doctor_additional_cities
(
    id        bigserial primary key,
    doctor_id bigint not null,
    city_id   bigint not null
);

create table if not exists docstar_site_doctor_additional_specialties
(
    id             bigserial primary key,
    doctor_id      bigint not null,
    speciallity_id bigint not null
);

create table if not exists docstar_site_doctor
(
    id                 bigserial primary key,
    name               varchar(100) not null,
    slug               varchar(100) not null,
    email              varchar(100) not null,
    inst_url           varchar(100),
    vk_url             varchar(100),
    dzen_url           varchar(100),
    tg_url             varchar(100),
    medical_directions varchar(255),
    main_blog_theme    text,
    prodoctorov        varchar(100),
    youtube_url        varchar(100),
    tg_channel_url     varchar(100),
    tiktok_url         varchar(100),

    is_active          bool,

    city_id            bigint,
    speciallity_id     bigint,

    date_created       timestamp    not null default now(),
    birth_date         date,

    s3_image           varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists docstar_site_city;
drop table if exists docstar_site_speciallity;
drop table if exists docstar_site_doctor_additional_cities;
drop table if exists docstar_site_doctor_additional_specialties;
drop table if exists docstar_site_doctor;
-- +goose StatementEnd
