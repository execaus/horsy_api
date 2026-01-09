-- +goose Up
-- +goose StatementBegin
create table if not exists app.breeds(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);
comment on table app.breeds is 'Справочник пород лошадей';
comment on column app.breeds.name is 'Название породы';
comment on column app.breeds.description is 'Описание породы';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists app.breeds cascade;
-- +goose StatementEnd