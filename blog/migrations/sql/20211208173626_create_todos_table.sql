-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    is_completed BOOLEAN DEFAULT false,
    
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todos ;
-- +goose StatementEnd
