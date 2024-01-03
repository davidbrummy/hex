package main

import (
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"

	gRPC "hex/internal/adapters/framework/left/grpc"

	"log"
	"os"
)

func main() {
	var err error

	var dbaseAdapter ports.DbPort
	var core ports.ArthimeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)

	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	// Domain Layer
	core = arithmetic.NewAdapter()

	// Application Layer
	appAdapter = api.NewAdapter(dbaseAdapter, core)

	// Framework Layer
	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()

}
