package main

import (

	// Import the generated protobuf code

	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/jobin212/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/jobin212/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
	)

	// Init will parse the command line flags
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentColelction := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentColelction}
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", srv.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
