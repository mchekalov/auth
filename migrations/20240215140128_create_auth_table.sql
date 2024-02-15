-- +goose Up
-- +goose StatementBegin
create table tuser (
    id serial primary key,
    uname text not null,
    email text not null,
    passwordhash varchar,
    urole integer,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table tuser;
-- +goose StatementEnd
