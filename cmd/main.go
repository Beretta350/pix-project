package main

import (
	"os"
	"pix-project/application/grpc"
	"pix-project/infra/db"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
