package main

import (
    pb "github.com/arturoguerra/destinyarena-faceit/proto"
)

func (f *FaceitServer) GetUserHubs(req *pb.ProfileRequest, stream pb.Faceit_GetUserHubsServer) error {
    log.Infof("Gettin hubs for: %s", req.GetGuid())
    hubs, err := f.API.GetUserHubs(req.GetGuid())
    if err != nil {
        log.Error(err)
        return err
    }

    for _, hub := range hubs {
        shub := &pb.Hub{
            Hubid:  hub.Hubid,
            Name:   hub.Name,
            Gameid: hub.GameID,
        }

        if err := stream.Send(shub); err != nil {
            log.Error(err)
            return err
        }
    }

    return nil
}
