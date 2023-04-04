package routes

// This is an example of how to handle a POST request and sending a gRPC message.

import (
	"context"
	"net/http"

	"github.com/Gemba-Kaizen/menumeister-api/pkg/ping/pb"
	"github.com/gin-gonic/gin"
)

// Expected HTTP request body
type CreatePingRequestBody struct {
	Message string `json:"message"`
}

// Receiving from POST request example
// HTTP handler that handles HTTP request
// Receive request and then make a gRPC call to microservice
func CreatePing(ctx *gin.Context, c pb.PingServiceClient) {
	// build the request body
	body := CreatePingRequestBody{}

	// parse the request body, throw error to requester if any issues
	// bind the POST request body to the fields defined in the struct
	if err := ctx.BindJSON(&body); err!= nil {
    ctx.AbortWithError(http.StatusBadRequest, err)
    return
  }

	// Create the gRPC message, re is a pointer to the variable
	res, err := c.CreatePing(context.Background(), &pb.CreatePingRequest{
    Message: body.Message,
	})

	if err!= nil {
    ctx.AbortWithError(http.StatusBadGateway, err)
    return
  }

	// respond to the request with success and the address of the ping request
	ctx.JSON(http.StatusCreated, &res)
}
