package main

import (
    "context"
    pb "github.com/arturoguerra/destinyarena-faceit/proto"
)

func (f *FaceitServer) Unban(ctx context.Context, in *pb.UnbanRequest) (*pb.UnbanReply, error) {
    err := f.API.Ban(in.GetHubid(), in.GetGuid(), in.GetReason())
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return &pb.UnbanReply{
        Status: "success",
    }, nil
}
