-- +goose Up
-- +goose StatementBegin
create table if not exists app.genders(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);

comment on table app.genders is 'Справочник полов лошадей';
comment on column app.genders.name is 'Название пола лошади';
comment on column app.genders.description is 'Описание пола лошади';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists app.genders cascade;
-- +goose StatementEnd