package db

// TYPE is a custom type to define kind of database
type TYPE string

// Constants
const (
	mysql    TYPE = "mysql"
	postgres TYPE = "postgres"
)

// NewPersistenceLayer return struct that implement DatabaseHandler
func NewPersistenceLayer(database TYPE, url string) (DatabaseHandler, error) {
	switch database {
	case mysql:
		return NewSQLLayer(string(database), url)
	case postgres:
		return NewSQLLayer(string(database), url)
	default:
		return nil, nil
	}
}
