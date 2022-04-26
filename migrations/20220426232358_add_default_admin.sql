-- +goose Up
-- +goose StatementBegin
INSERT INTO rools (name)
                  VALUES ('admin');

INSERT INTO rools (name)
VALUES ('user');

INSERT INTO users (name, telegram_id)
VALUES ( 'MrS','519588080');

INSERT INTO user_rools (user_id, rool_id)
VALUES ('1', '1');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM rools
    WHERE name = 'admin';

DELETE FROM rools
    WHERE name = 'user';

DELETE FROM users
    WHERE name = 'MrS' and telegram_id = '519588080';

DELETE FROM user_rools
    WHERE user_id = '1' and rool_id = '1';
-- +goose StatementEnd
