package main

import (
    "context"
    pb "github.com/arturoguerra/destinyarena-faceit/proto"
)

func (f *FaceitServer) GetInvite(ctx context.Context, in *pb.InviteRequest) (*pb.InviteReply, error) {
    err, code := f.API.GetInvite(in.GetHubid())
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return &pb.InviteReply{
        Base: "https://www.faceit.com/en/inv",
        Code: code,
    }, nil
}
