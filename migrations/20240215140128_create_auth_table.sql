-- +goose Up
-- +goose StatementBegin
create table auth_user (
    id serial primary key,
    user_name text not null,
    email text not null,
    passwordhash text,
    user_role integer,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table auth_user;
-- +goose StatementEnd
