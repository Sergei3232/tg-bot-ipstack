-- +goose Up
-- +goose StatementBegin
insert into rols (name)
values ('admin');

insert into rols (name)
values ('user');

insert into users (name, telegram_id)
values ( 'MrS','519588080');

insert into user_rols (user_id, rol_id)
values ('1', '1');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from rols
    where name = 'admin';

delete from rols
    where name = 'user';

delete from users
    where name = 'MrS' and telegram_id = '519588080';

delete from user_rols
    where user_id = '1' and rol_id = '1';
-- +goose StatementEnd
