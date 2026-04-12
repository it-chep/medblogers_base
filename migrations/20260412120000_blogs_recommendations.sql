-- +goose Up
-- +goose StatementBegin
create table if not exists blogs_recommendations
(
    id                     bigserial primary key,
    blog_id                uuid   not null,
    recommended_blog_id    uuid   not null,

    constraint uq_blogs_recommendations unique (blog_id, recommended_blog_id),
    constraint chk_blogs_recommendations_not_self check (blog_id <> recommended_blog_id),
    constraint fk_blogs_recommendations_blog foreign key (blog_id) references blog (id) on delete cascade,
    constraint fk_blogs_recommendations_recommended_blog foreign key (recommended_blog_id) references blog (id) on delete cascade
);

create index if not exists idx_blogs_recommendations_blog_id on blogs_recommendations (blog_id);
create index if not exists idx_blogs_recommendations_recommended_blog_id on blogs_recommendations (recommended_blog_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists blogs_recommendations;
-- +goose StatementEnd
