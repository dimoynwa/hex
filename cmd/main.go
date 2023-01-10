package main

import (
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	rpc "hex/internal/adapters/framework/left/grpc"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"
	"log"
	"os"
)

func main() {
	var err error

	// Ports
	var dbAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	dbDriver := os.Getenv("DB_DRIVER") // "postgres"
	dsName := os.Getenv("DS_NAME")     // "jdbc:postgresql://localhost:5432/postgres"

	dbAdapter, err = db.NewAdapter(dbDriver, dsName)
	if err != nil {
		log.Fatalf("failed to start db: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, core)

	grpcAdapter = rpc.NewAdapter(appAdapter)
	grpcAdapter.Run()
}
