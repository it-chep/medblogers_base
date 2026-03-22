-- +goose Up
create table if not exists banners
(
    id                bigserial primary key,
    is_active         bool default false,
    name              text,

    desktop_image     uuid,
    desktop_file_type text,

    mobile_image      uuid,
    mobile_file_type  text,

    banner_link       text,
    ordering_number   bigint default 999
);

-- +goose Down
drop table if exists banners;
