package routes

// This is an example of how to handle a GET request and send a gRPC message

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Gemba-Kaizen/menumeister-api/pkg/ping/pb"
	"github.com/gin-gonic/gin"
)

// Get request, no need of binding of request body
func GetPing(ctx *gin.Context, c pb.PingServiceClient) {
	// parse url params and get string as it is
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	// Create the gRPC message, res is a pointer to the variable
	res, err := c.GetPing(context.Background(), &pb.GetPingRequest{
		UserId: int64(id),
	})

	if err!= nil {
    ctx.AbortWithError(http.StatusBadGateway, err)
    return
  }

	// respond to the request with success and the address of the ping request
	ctx.JSON(http.StatusCreated, &res)
}
