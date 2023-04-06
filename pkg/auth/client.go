package auth

import (
	"fmt"

	"github.com/Gemba-Kaizen/menumeister-api/config"
	"github.com/Gemba-Kaizen/menumeister-api/pkg/auth/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// Insecure, no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err!= nil {
    fmt.Println("Could not connect: ", err)
  }

	return pb.NewAuthServiceClient(cc)
}