module github.com/jorgeAM/events/bookingService

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.12
	github.com/joho/godotenv v1.3.0
	github.com/jorgeAM/events/msgbroker v0.0.0
	github.com/lib/pq v1.5.2 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pressly/goose v2.6.0+incompatible
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/ziutek/mymysql v1.5.4 // indirect
	google.golang.org/appengine v1.6.6 // indirect
)

replace github.com/jorgeAM/events/msgbroker v0.0.0 => ../msgbroker
