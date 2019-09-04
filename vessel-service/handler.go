package main

import (
	"context"

	pb "github.com/jobin212/shippy/vessel-service/proto/vessel"
)

type handler struct {
	repository
}

func (s *handler) FindAvailable(ctx context.Context,
	req *pb.Specification,
	res *pb.Response) error {

	vessel, err := s.repository.FindAvailable(req)

	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repository.Create(req); err != nil {
		return err
	}

	res.Vessel = req
	res.Created = true
	return nil
}
