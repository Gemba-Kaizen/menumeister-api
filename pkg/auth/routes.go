package auth

import (
	"github.com/Gemba-Kaizen/menumeister-api/config"
	"github.com/Gemba-Kaizen/menumeister-api/pkg/auth/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient{
	// If you have any middlewares, you can init them here

	// Initialise gRPC client service
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	// HTTP endpoints
	routes := r.Group("/auth")
	routes.POST("/register", svc.RegisterMerchant)
	routes.POST("/login", svc.LoginMerchant)

	return svc
}

// functions defined for register route to connect handlers
func (svc *ServiceClient) RegisterMerchant(ctx *gin.Context) {
	// CreatePing function in routes package
	routes.RegisterMerchant(ctx, svc.Client)
}

func (svc *ServiceClient) LoginMerchant(ctx *gin.Context) {
	routes.LoginMerchant(ctx, svc.Client)
}