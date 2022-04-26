-- +goose Up
-- +goose StatementBegin
create table rools
(
    id  SERIAL primary key,
    name varchar not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists rools;
-- +goose StatementEnd
