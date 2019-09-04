package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/jobin212/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

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

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("NO MONGO")
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	log.Println("CONNECTED TO MONGO")

	vesselsCollection := client.Database("shippy").Collection("vessels")

	repository := &VesselRepository{vesselsCollection}

	// Register implementation
	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repository})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
