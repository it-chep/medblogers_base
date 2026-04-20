-- +goose Up
alter table freelancer
    add column if not exists has_med_education boolean not null default false;

alter table freelancers_price_list
    add column if not exists price_to integer,
    add column if not exists search_vector tsvector;

alter table freelancers_price_list
    drop constraint if exists freelancers_price_list_freelancer_id_name_price_key;

update freelancers_price_list
set search_vector = to_tsvector('russian', coalesce(name, ''))
where search_vector is null;

create index if not exists idx_freelancers_price_list_search_vector
    on freelancers_price_list
    using gin (search_vector);

-- +goose Down
drop index if exists idx_freelancers_price_list_search_vector;

alter table freelancers_price_list
    drop column if exists search_vector,
    drop column if exists price_to;

alter table freelancers_price_list
    add constraint freelancers_price_list_freelancer_id_name_price_key
        unique (freelancer_id, name, price);

alter table freelancer
    drop column if exists has_med_education;
