package main

import (
	"context"

	pb "github.com/destinyarena/faceit/proto"
)

func (f *faceitService) GetProfileByID(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	log.Infoln(in.GetGuid())
	u, err := f.API.GetUserByID(in.GetGuid())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.ProfileReply{
		Guid:     u.ID,
		Username: u.Username,
		Skilllvl: int32(u.SkillLevel),
	}, nil
}

func (f *faceitService) GetProfileByName(ctx context.Context, in *pb.ProfileNameRequest) (*pb.ProfileReply, error) {
	log.Infoln(in.GetName())
	u, err := f.API.GetUserByName(in.GetName())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.ProfileReply{
		Guid:     u.ID,
		Username: u.Username,
		Skilllvl: int32(u.SkillLevel),
	}, nil
}
