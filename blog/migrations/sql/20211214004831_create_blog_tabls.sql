-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blogs(
    id SERIAL NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    is_completed BOOLEAN DEFAULT false,

    PRIMARY KEY(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS blogs;
-- +goose StatementEnd
