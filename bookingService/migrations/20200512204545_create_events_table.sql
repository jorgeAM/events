-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE events (
    id varchar(30) NOT NULL,
    name varchar(100),
    start timestamp,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE events;
