package main

import (
	"context"

	pb "github.com/destinyarena/faceit/proto"
)

func (f *faceitService) Unban(ctx context.Context, in *pb.UnbanRequest) (*pb.Empty, error) {
	err := f.API.Unban(in.GetHubid(), in.GetGuid())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.Empty{}, nil
}
