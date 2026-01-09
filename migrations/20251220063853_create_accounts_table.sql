-- +goose Up
-- +goose StatementBegin
create table if not exists app.accounts(
    id uuid primary key, -- Уникальный идентификатор аккаунта
    email varchar(255) not null unique, -- Адрес электронной почты пользователя
    password varchar(255) not null, -- Хэш пароля пользователя
    created_at timestamp not null default now(), -- Дата и время создания аккаунта
    updated_at timestamp not null default now(), -- Дата и время последнего обновления аккаунта
    last_activity_at timestamp not null default now() -- Дата и время последней активности пользователя
);

comment on table app.accounts is 'Таблица хранения аккаунтов пользователей';

comment on column app.accounts.id is 'Уникальный идентификатор аккаунта';
comment on column app.accounts.email is 'Адрес электронной почты пользователя';
comment on column app.accounts.password is 'Хэш пароля пользователя';
comment on column app.accounts.created_at is 'Дата и время создания аккаунта';
comment on column app.accounts.updated_at is 'Дата и время последнего обновления аккаунта';
comment on column app.accounts.last_activity_at is 'Дата и время последней активности пользователя';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists app.accounts;
-- +goose StatementEnd