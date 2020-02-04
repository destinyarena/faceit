package main

import (
    "context"
    pb "github.com/arturoguerra/destinyarena-faceit/proto"
)

func (f *FaceitServer) Ban(ctx context.Context, in *pb.BanRequest) (*pb.Empty, error) {
    err := f.API.Ban(in.GetHubid(), in.GetGuid(), in.GetReason())
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return &pb.Empty{}, nil
}
