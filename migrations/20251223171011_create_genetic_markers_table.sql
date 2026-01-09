-- +goose Up
-- +goose StatementBegin
create table if not exists app.genetic_markers(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);

comment on table app.genetic_markers is 'Справочник генетических маркеров';
comment on column app.genetic_markers.name is 'Название генетического маркера';
comment on column app.genetic_markers.description is 'Описание генетического маркера';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists app.genetic_markers cascade;
-- +goose StatementEnd