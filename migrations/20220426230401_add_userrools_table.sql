-- +goose Up
-- +goose StatementBegin
create table user_rools
(
    id  SERIAL primary key,
    user_id integer not null,
    rool_id integer not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_rools;
-- +goose StatementEnd
