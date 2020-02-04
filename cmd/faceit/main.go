package main

import (
    "fmt"
    "net"
    "google.golang.org/grpc"
    pb "github.com/arturoguerra/destinyarena-faceit/proto"
    faceit "github.com/arturoguerra/destinyarena-faceit/internal/restapi"
    "github.com/arturoguerra/destinyarena-faceit/internal/config"
    "github.com/arturoguerra/destinyarena-faceit/internal/logging"
)

var log = logging.New()

type FaceitServer struct {
    API *faceit.Faceit
    pb.UnimplementedFaceitServer
}

func main() {
    cfg := config.LoadConfig()
    host := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
    lis, err := net.Listen("tcp", host)
    if err != nil {
        log.Fatalf(err.Error())
    }
    log.Infof("Listening on: %s", host)

    s := grpc.NewServer()

    api := faceit.New(cfg.ApiToken, cfg.UserToken)

    fs := &FaceitServer{
        API: api,
    }
    pb.RegisterFaceitServer(s, fs)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
