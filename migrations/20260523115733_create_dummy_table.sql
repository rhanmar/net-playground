-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS dummy (
    id SERIAL PRIMARY KEY,
    data VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dummy;
-- +goose StatementEnd
