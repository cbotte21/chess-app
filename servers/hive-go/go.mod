module github.com/cbotte21/hive-go

go 1.19

require github.com/cbotte21/judicial-go v0.0.0

replace github.com/cbotte21/judicial-go v0.0.0 => ../judicial-go

require (
	github.com/cbotte21/microservice-common v1.0.0
	golang.org/x/net v0.4.0
	google.golang.org/grpc v1.52.0
	google.golang.org/protobuf v1.28.1
)

replace github.com/cbotte21/microservice-common => ../microservice-common

require (
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
)
