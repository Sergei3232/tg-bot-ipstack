-- +goose Up
-- +goose StatementBegin
INSERT INTO rols (name)
VALUES ('admin');

INSERT INTO rols (name)
VALUES ('user');

INSERT INTO users (name, telegram_id)
VALUES ( 'MrS','519588080');

INSERT INTO user_rols (user_id, rol_id)
VALUES ('1', '1');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM rols
    WHERE name = 'admin';

DELETE FROM rols
    WHERE name = 'user';

DELETE FROM users
    WHERE name = 'MrS' and telegram_id = '519588080';

DELETE FROM user_rols
    WHERE user_id = '1' and rol_id = '1';
-- +goose StatementEnd
