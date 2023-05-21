package auth

import (
	"carshop/internal/model"
	"carshop/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	AccessTokenHeader = "Auth-Access-Token"
)

func GetIssuer(ctx *gin.Context) (role model.Role, token string) {

	logger.Log.Info("validating token...")

	receivedToken := ctx.Request.Header.Get(AccessTokenHeader)

	if receivedToken == "" {
		logger.Log.Error("login failed, token is empty")
		return
	}

	for _, validToken := range Tokens {

		if validToken.AccessToken == receivedToken {
			logger.Log.Info("access granted",
				zap.String("role", string(validToken.Role)),
			)

			role = model.Role(validToken.Role)
			token = validToken.AccessToken
			return
		}
	}

	return "", ""
}

func CheckRole(ctx *gin.Context, requiredRole model.Role) bool {
	role, _ := GetIssuer(ctx)

	return role == requiredRole
}
