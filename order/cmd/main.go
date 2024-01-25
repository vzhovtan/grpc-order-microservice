package main

import (
	"log"
	"microservices/order/config"
	"microservices/order/internal/adapters/db"
	"microservices/order/internal/adapters/grpc"
	"microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSrc())
	if err != nil {
		log.Fatalf("Failed to connet to DB. Error is: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, 3000)
	grpcAdapter.Run()
}
