package routes

import (
	"context"
	"net/http"

	"github.com/Gemba-Kaizen/menumeister-api/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type MerchantCredentialFields struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func RegisterMerchant(ctx *gin.Context, c pb.AuthServiceClient) {
	body := MerchantCredentialFields{}

	if err := ctx.ShouldBindJSON(&body); err!= nil {
    ctx.AbortWithError(http.StatusBadRequest, err)
    return
  }

	res, err := c.RegisterMerchant(context.Background(), &pb.RegisterMerchantRequest{
		Email: body.Email,
		Password: body.Password,
	})

	if err != nil {
    ctx.AbortWithError(http.StatusBadGateway, err)
    return
  }

	if res.Error != "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": res.Error})
	}

	ctx.JSON(http.StatusCreated, &res)
}

func LoginMerchant(ctx *gin.Context, c pb.AuthServiceClient) {
	body := MerchantCredentialFields{}

	if err := ctx.ShouldBindJSON(&body); err!= nil {
    ctx.AbortWithError(http.StatusBadRequest, err)
    return
  }

	res, err := c.LoginMerchant(context.Background(), &pb.LoginMerchantRequest{
		Email: body.Email,
		Password: body.Password,
	})

	if err != nil {
    ctx.AbortWithError(http.StatusBadGateway, err)
    return
  }

	if res.Error != "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": res.Error})
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
