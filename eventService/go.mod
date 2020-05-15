module github.com/jorgeAM/events/eventService

go 1.13

require (
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/joho/godotenv v1.3.0
	github.com/jorgeAM/events/msgbroker v0.0.0
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
)

replace github.com/jorgeAM/events/msgbroker v0.0.0 => ../msgbroker
