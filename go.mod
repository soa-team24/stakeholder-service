module stakeholder-service

go 1.22.1

require (
	github.com/google/uuid v1.6.0
	gorm.io/driver/mysql v1.5.5
	gorm.io/gorm v1.25.8
	soa/grpc/proto v0.0.1
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240513163218-0867130af1f8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240509183442-62759503f434 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gorilla/handlers v1.5.2
	github.com/gorilla/mux v1.8.1
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	google.golang.org/grpc v1.64.0
)

replace soa/grpc/proto => /app/proto
