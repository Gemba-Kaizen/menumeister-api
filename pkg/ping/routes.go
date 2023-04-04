// This is where your endpoints are defined

package ping

import (
	"github.com/Gemba-Kaizen/menumeister-api/config"
	"github.com/Gemba-Kaizen/menumeister-api/pkg/ping/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) {
	// If you have any middlewares, you can init them here

	// Initialise gRPC client service
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	// HTTP endpoints
	routes := r.Group("/ping")
	routes.GET("/:id", svc.GetPing)
	routes.POST("/create", svc.CreatePing)
}

// functions defined for register route to connect handlers
func (svc *ServiceClient) CreatePing(ctx *gin.Context) {
	// CreatePing function in routes package
	routes.CreatePing(ctx, svc.Client)
}

func (svc *ServiceClient) GetPing(ctx *gin.Context) {
	routes.GetPing(ctx, svc.Client)
}