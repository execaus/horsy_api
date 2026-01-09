-- +goose Up
-- +goose StatementBegin
create table if not exists app.birthplaces(
    id uuid primary key default gen_random_uuid(),
    name varchar not null unique,
    description text
);

comment on table app.birthplaces is 'Справочник мест рождения лошадей';
comment on column app.birthplaces.name is 'Название места рождения';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists app.birthplaces cascade;
-- +goose StatementEnd