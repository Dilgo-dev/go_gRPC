package main

import (
	"fmt"
	"go_server/internal/api/routes"
	"go_server/internal/grpc"
	"go_server/internal/utils"
	"log"
	"net"
	"net/http"
	"os"

	chi "github.com/go-chi/chi"
	middleware "github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	grpc_go "google.golang.org/grpc"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Error loading .env file")
		}
	}

	utils.InitDatabase()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("PORT env not set, using default: 8080")
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
		log.Println("GRPC_PORT env not set, using default: 50051")
	}

	go startGrpcServer(grpcPort)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/logs", routes.LogRouter())

	fmt.Printf("Server go is listening on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, r)
}

func startGrpcServer(grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc_go.NewServer()
	grpc.RegisterGrpcServer(s)
	
	fmt.Printf("gRPC server listening on :%s\n", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}