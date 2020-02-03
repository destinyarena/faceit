package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "github.com/arturoguerra/destinyarena-faceit/proto"
)

const addr = "localhost:3000"

func main() {
    conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatal(err.Error())
    }

    defer conn.Close()

    c := pb.NewFaceitClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := c.GetInvite(ctx, &pb.InviteRequest{
        Hubid: "86f3b3ff-de0b-4d09-8b86-13f621dd32aa",
    })

    if err != nil {
        log.Fatal(err.Error())
    }

    log.Printf("%s/%s", r.GetBase(), r.GetCode())

    ctx, cancel = context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    e, err := c.GetProfile(ctx, &pb.ProfileRequest{
        Guid: "",
    })
    if err != nil {
        log.Fatal(err.Error())
    }

    log.Printf("%s %s %s", e.GetGuid(), e.GetSkilllvl, e.GetUsername())

}
