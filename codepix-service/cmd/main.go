package main

import (
	"os"

	"github.com/robertheory/codepix-go/app/grpc"
	"github.com/robertheory/codepix-go/infra/db"
)

func main() {

	database := db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)

}
