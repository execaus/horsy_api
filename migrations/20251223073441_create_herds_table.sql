-- +goose Up
-- +goose StatementBegin
create table if not exists app.herds(
    id uuid primary key, -- Уникальный идентификатор стада
    name varchar(64) not null, -- Название стада
    description text, -- Описание стада
    account_id uuid not null, -- Идентификатор аккаунта, владельца стада
    created_at timestamp not null default now(), -- Дата и время создания записи
    updated_at timestamp not null default now() -- Дата и время последнего обновления записи
);

comment on table app.herds is 'Таблица хранения данных о стадах';

comment on column app.herds.id is 'Уникальный идентификатор стада';
comment on column app.herds.name is 'Название стада';
comment on column app.herds.description is 'Описание стада';
comment on column app.herds.account_id is 'Идентификатор аккаунта, владельца стада';
comment on column app.herds.created_at is 'Дата и время создания';
comment on column app.herds.updated_at is 'Дата и время последнего обновления';

alter table app.herds add constraint fk_herds_account foreign key (account_id) references app.accounts(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists app.herds;
-- +goose StatementEnd