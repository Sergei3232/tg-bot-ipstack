-- +goose Up
-- +goose StatementBegin
create table rols
(
    id  SERIAL primary key,
    name varchar not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists rols;
-- +goose StatementEnd
