-- +goose Up
-- +goose StatementBegin
create table users
(
    id SERIAL primary key,
    name varchar not null,
    telegram_id integer not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
