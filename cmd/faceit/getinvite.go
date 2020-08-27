package main

import (
	"context"

	pb "github.com/destinyarena/faceit/proto"
)

func (f *faceitService) GetInvite(ctx context.Context, in *pb.InviteRequest) (*pb.InviteReply, error) {
	code, err := f.API.GetInvite(in.GetHubid())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.InviteReply{
		Base: "https://www.faceit.com/en/inv",
		Code: code,
	}, nil
}
