package main

import (
	"fmt"
	"net"

	"github.com/destinyarena/faceit/internal/config"
	"github.com/destinyarena/faceit/internal/logging"
	faceit "github.com/destinyarena/faceit/internal/restapi"
	pb "github.com/destinyarena/faceit/proto"
	"google.golang.org/grpc"
)

var log = logging.New()

type faceitService struct {
	API faceit.Faceit
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

	api, err := faceit.New(cfg.ApiToken, cfg.UserToken)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fs := &faceitService{
		API: api,
	}

	pb.RegisterFaceitServer(s, fs)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
