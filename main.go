package main

import (
	"log"
	"os"
	"os/signal"
	"stakeholder-service/handler"
	"stakeholder-service/model"
	"stakeholder-service/repository"
	"stakeholder-service/service"
	"syscall"

	"net"

	"stakeholder-service/proto/stakeholder"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "root:root@tcp(stakeholder_db:3307)/soa?charset=utf8mb4&parseTime=True&loc=Local"
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

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	stakeholder.RegisterStakeholderServiceServer(grpcServer, authHandler)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
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

	print("Stakeholder starting print")
	log.Println("Stake holde staring")
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
