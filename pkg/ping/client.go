package ping

import (
	"fmt"

	"github.com/Gemba-Kaizen/menumeister-api/config"
	"github.com/Gemba-Kaizen/menumeister-api/pkg/ping/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.PingServiceClient
}

func InitServiceClient(c *config.Config) pb.PingServiceClient {
	// Insecure, no SSL running
	cc, err := grpc.Dial(c.PingSvcUrl, grpc.WithInsecure())

	if err!= nil {
    fmt.Println("Could not connect: ", err)
  }

	return pb.NewPingServiceClient(cc)
}