-- +goose Up
-- +goose StatementBegin
create table user_rols
(
    id  SERIAL primary key,
    user_id integer not null,
    rol_id integer not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_rols;
-- +goose StatementEnd
