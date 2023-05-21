package auth

import (
	"carshop/internal/model"
	"carshop/internal/request"
	"carshop/internal/route/fail"
	"carshop/internal/server/router"
	"carshop/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logout(ctx *gin.Context) {

	var (
		logoutRequest  = request.Request{}
		logoutResponse = request.Response{}
	)

	log := logger.Log.WithOptions(zap.Fields(
		zap.Any("data", logoutRequest.Data),
	))

	// Parse the posted JSON data.
	err := ctx.ShouldBind(&logoutRequest)
	if err != nil {
		logger.Log.Error("Failed to bind input data",
			zap.Error(err),
		)

		fail.ReturnError(ctx, logoutResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	log.Info("logout started")

	// check authentication
	role, token := GetIssuer(ctx)
	if role == model.Empty || token == "" {
		logger.Log.Error("failed to authorize user",
			zap.Error(err),
		)

		fail.ReturnError(ctx, logoutResponse, []string{"failed to authorize user"}, 400, logger.Log)
		return
	}

	tempTokens := make([]model.Token, len(Tokens))
	for _, nextToken := range Tokens {
		if nextToken.AccessToken == token {
			continue
		}
		tempTokens = append(tempTokens, nextToken)
	}

	copy(Tokens, tempTokens)

	ctx.Header(AccessTokenHeader, "")

	log.Info("logout finished")

	logoutResponse.Status = true
	ctx.JSON(200, logoutResponse)
}

func init() {
	router.Router.Handle("POST", "auth/logout", Logout)
}
