package main

import (
	"context"
	"log"

	pb "github.com/jobin212/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/jobin212/shippy/vessel-service/proto/vessel"
)

type handler struct {
	repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	// Set the VesselId as the vessel we got back from our service
	req.VesselId = vesselResponse.Vessel.Id

	if err = s.repository.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repository.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
