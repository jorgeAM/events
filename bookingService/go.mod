module github.com/jorgeAM/events/bookingService

go 1.13

require (
	github.com/gorilla/mux v1.7.4
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/jorgeAM/events/msgbroker v0.0.0 // indirect
)

replace github.com/jorgeAM/events/msgbroker v0.0.0 => ../msgbroker
