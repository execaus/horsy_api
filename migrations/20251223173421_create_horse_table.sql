-- +goose Up
-- +goose StatementBegin
-- Основная таблица лошадей
create table if not exists app.horses(
    id uuid primary key default gen_random_uuid(),
    herd uuid not null,
    gender uuid,
    name varchar,
    birth_day int,
    birth_month int,
    birth_year int,
    birth_place uuid,
    withers_height int,
    sire uuid,
    dam uuid,
    is_pregnant boolean not null,
    is_dead boolean not null,
    description text,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null,
    constraint uq_name_herd unique (name, herd)
);
comment on table app.horses is 'Основная таблица с информацией о лошадях';

comment on column app.horses.id is 'Уникальный идентификатор лошади';
comment on column app.horses.herd is 'Идентификатор табуна, к которому принадлежит лошадь';
comment on column app.horses.gender is 'Пол лошади';
comment on column app.horses.name is 'Имя лошади';
comment on column app.horses.birth_day is 'День рождения лошади';
comment on column app.horses.birth_month is 'Месяц рождения лошади';
comment on column app.horses.birth_year is 'Год рождения лошади';
comment on column app.horses.birth_place is 'Место рождения лошади';
comment on column app.horses.withers_height is 'Высота в холке (в сантиметрах)';
comment on column app.horses.sire is 'Идентификатор отца (жеребца)';
comment on column app.horses.dam is 'Идентификатор матери (кобылы)';
comment on column app.horses.is_pregnant is 'Статус беременности (true, если беременна)';
comment on column app.horses.is_dead is 'Статус лошади (true, если мертва)';
comment on column app.horses.created_at is 'Время создания записи';
comment on column app.horses.updated_at is 'Время последнего обновления записи';
comment on column app.horses.description is 'Описание лошади';

alter table app.horses add constraint fk_recursive_sire foreign key (sire) references app.horses(id) on delete set null;
alter table app.horses add constraint fk_recursive_dam foreign key (dam) references app.horses(id) on delete set null;
alter table app.horses add constraint fk_gender foreign key (gender) references app.genders(id);
alter table app.horses add constraint fk_birth_place foreign key (birth_place) references app.birthplaces(id);

-- Связующие таблицы
create table if not exists app.horse_color(
    horse uuid not null,
    color uuid not null,
    constraint pk_horse_color primary key (horse, color)
);
comment on table app.horse_color is 'Таблица для хранения окрасов лошадей';

comment on column app.horse_color.horse is 'Идентификатор лошади';
comment on column app.horse_color.color is 'Идентификатор окраса';

alter table app.horse_color add constraint fk_horse_colors_horse foreign key (horse) references app.horses(id) on delete cascade;
alter table app.horse_color add constraint fk_horse_colors_color foreign key (color) references app.colors(id);

create table if not exists app.horse_genetic_marker(
    horse uuid not null,
    marker uuid not null,
    constraint pk_horse_genetic_marker primary key (horse, marker)
);
comment on table app.horse_genetic_marker is 'Таблица для хранения генетических маркеров лошадей';

comment on column app.horse_genetic_marker.horse is 'Идентификатор лошади';
comment on column app.horse_genetic_marker.marker is 'Идентификатор генетического маркера';

alter table app.horse_genetic_marker add constraint fk_horse_genetic_marker_horse foreign key (horse) references app.horses(id) on delete cascade;
alter table app.horse_genetic_marker add constraint fk_horse_genetic_marker_marker foreign key (marker) references app.genetic_markers(id);

create table if not exists app.horse_breed(
    horse uuid not null,
    breed uuid not null,
    percent int not null,
    constraint pk_horse_breed primary key (horse, breed),
    constraint ck_breed_percent check (percent >= 0 and percent <= 10000)
);
comment on table app.horse_breed is 'Таблица для хранения породной принадлежности лошадей с указанием процента';

comment on column app.horse_breed.horse is 'Идентификатор лошади';
comment on column app.horse_breed.breed is 'Идентификатор породы';
comment on column app.horse_breed.percent is 'Процент принадлежности к породе, целое число от 0 до 10000, где 10000 = 100%';

alter table app.horse_breed add constraint fk_horse_breeds_horse foreign key (horse) references app.horses(id) on delete cascade;
alter table app.horse_breed add constraint fk_horse_breeds_breed foreign key (breed) references app.breeds(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists app.horse_breed cascade;
drop table if exists app.horse_genetic_marker cascade;
drop table if exists app.horse_color cascade;
drop table if exists app.horses cascade;
-- +goose StatementEnd