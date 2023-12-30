-- +goose Up
-- +goose StatementBegin

create table dir(
    id uuid not null primary key,
    parent_path varchar(100) not null,

    constraint empty_dir_parent check(parent_path != '')
);

create table file(
    id uuid not null primary key,
    key varchar(100) not null unique,
    parent_path varchar(100) not null,

    constraint empty_file_parent_path check(parent_path != ''),
    constraint empty_file_key check(key != '')
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table file;
drop table dir;

-- +goose StatementEnd