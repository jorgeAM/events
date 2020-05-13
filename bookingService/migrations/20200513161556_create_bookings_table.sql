-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE bookings (
  id int NOT NULL,
	start timestamp NOT NULL,
	seats int NOT NULL,
	eventId varchar(30) NOT NULL,
	createdAt timestamp NOT NULL,
	updatedAt timestamp NOT NULL,
  PRIMARY KEY(id)
);

ALTER TABLE bookings ADD CONSTRAINT FK_EventBooking FOREIGN KEY (eventId) REFERENCES events(id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE bookings;