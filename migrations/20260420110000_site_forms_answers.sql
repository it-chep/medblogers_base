-- +goose Up
create table if not exists site_forms_answers
(
    id         bigserial primary key,
    form_name  text      not null,
    answer     jsonb     not null,
    created_at timestamp not null default now(),
    cookie_id  uuid      not null,
    source     text      not null,
    tg         text
);

create index if not exists idx_site_forms_answers_form_name on site_forms_answers (form_name);
create index if not exists idx_site_forms_answers_form_name on site_forms_answers (source);
create index if not exists idx_site_forms_answers_created_at on site_forms_answers (created_at desc);

-- +goose Down
drop table if exists site_forms_answers;
