-- +goose Up
-- +goose StatementBegin
create table if not exists app.colors(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);

comment on table app.colors is 'Справочник окрасов лошадей';
comment on column app.colors.name is 'Название окраса лошади';
comment on column app.colors.description is 'Описание окраса лошади';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists app.colors cascade;
-- +goose StatementEnd