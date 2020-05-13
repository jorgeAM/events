-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE bookings (
  id SERIAL NOT NULL,
	start timestamp NOT NULL,
	seats int NOT NULL,
	event_id varchar(30) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
  PRIMARY KEY(id)
);

ALTER TABLE bookings ADD CONSTRAINT FK_EventBooking FOREIGN KEY (event_id) REFERENCES events(id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE bookings;