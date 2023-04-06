package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/Gemba-Kaizen/menumeister-api/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
    ctx.AbortWithStatus(http.StatusUnauthorized)
    return
  }

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
    return
	}

	res, err := c.svc.Client.ValidateMerchant(context.Background(), &pb.ValidateMerchantRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
    ctx.AbortWithStatus(http.StatusUnauthorized)
    return
  }

	ctx.Set("merchatId", res.MerchantId)
	ctx.Next()
}