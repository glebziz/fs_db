-- +goose Up
-- +goose StatementBegin

alter table file rename to file_old;

create table content_file(
    id uuid not null primary key,
    parent_path varchar(512) not null,

    constraint empty_file_parent_path check(parent_path != '')
);

create table file(
    id integer primary key autoincrement,
    key varchar(100) not null,
    content_id uuid references content_file(id),
    tx_id uuid not null default '00000000-0000-0000-0000-000000000000',
    ts datetime not null default (datetime('now')),

    constraint empty_file_key check(key != '')
);

insert into content_file(id, parent_path)
select id, parent_path from file_old;

insert into file(key, content_id)
select key, id from file_old;

drop table file_old;

pragma foreign_keys = 1;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

alter table file rename to file_old;

create table file(
    id uuid not null primary key,
    key varchar(100) not null unique,
    parent_path varchar(100) not null,

    constraint empty_file_parent_path check(parent_path != ''),
    constraint empty_file_key check(key != '')
);

insert into file(id, key, parent_path)
with ranked_files as (
    select cf.id, fo.key, cf.parent_path,
           row_number() over (partition by fo.key order by fo.ts desc) old_tag
    from file_old fo
             inner join content_file cf on fo.content_id = cf.id
    where fo.tx_id = '00000000-0000-0000-0000-000000000000'
)
select id, key, parent_path
from ranked_files
where old_tag = 1;

drop table content_file;
drop table file_old;

-- +goose StatementEnd