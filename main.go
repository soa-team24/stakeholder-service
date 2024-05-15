package main

import (
	"log"
	"stakeholder-service/handler"
	"stakeholder-service/model"
	"stakeholder-service/repository"
	"stakeholder-service/service"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"soa/grpc/proto/stakeholder"
)

func initDB() *gorm.DB {
	connectionStr := "root:super@tcp(localhost:3306)/soa?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.User{}, &model.Person{}, &model.Profile{})

	return database
}
func startServer(authHandler *handler.AuthenticationHandler) {
	//router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc("/", authHandler.RegisterTourist).Methods("POST")
	//router.HandleFunc("/login", authHandler.Login).Methods("POST")

	lis, err := net.Listen("tcp", "localhost:8085")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	stakeholder.RegisterStakeholdersServiceServer(grpcServer, authHandler)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)

	/*
		allowedOrigins := handlers.AllowedOrigins([]string{"*"})
		allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
		allowedHeaders := handlers.AllowedHeaders([]string{
			"Content-Type",
			"Authorization",
			"X-Custom-Header",
		})*/

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	//corsRouter := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)

	println("Server starting")
	//log.Fatal(http.ListenAndServe(":8085", corsRouter))

}

func main() {

	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	userRepository := &repository.UserRepository{DatabaseConnection: database}
	profileRepository := &repository.ProfileRepository{DatabaseConnection: database}
	jwtGenerator := repository.NewJwtGenerator()

	authService := service.NewAuthenticationService(userRepository, jwtGenerator, profileRepository)
	authHandler := &handler.AuthenticationHandler{AuthenticationService: authService}

	startServer(authHandler)

}
