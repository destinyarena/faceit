package main

import (
    "context"
    pb "github.com/arturoguerra/destinyarena-faceit/proto"
)

func (f *FaceitServer) GetProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
    u, err := f.API.GetUser(in.GetGuid())
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return &pb.ProfileReply{
        Guid: u.Id,
        Username: u.Username,
        Skilllvl: int32(u.SkillLevel),
    }, nil
}
