package auth

import (
	"carshop/internal/model"
	"carshop/internal/request"
	"carshop/internal/route/fail"
	"carshop/internal/server/router"
	"carshop/pkg/crypto"
	"carshop/pkg/logger"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(ctx *gin.Context) {
	var (
		loginRequest                    = request.Request{}
		loginResponse                   = request.Response{}
		username, password, token, role string
		ok                              bool
	)

	// Parse the posted JSON data.
	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		logger.Log.Error("Failed to bind input data",
			zap.Error(err),
		)

		fail.ReturnError(ctx, loginResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	log := logger.Log.WithOptions(zap.Fields(
		zap.Any("data", loginRequest.Data),
	))

	log.Info("login started")

	if username, ok = loginRequest.Data["username"].(string); !ok {
		err := fmt.Errorf("username is missing")
		log.Error("auth/login failed")

		fail.ReturnError(ctx, loginResponse, []string{err.Error()}, 400, log)
		return
	}

	if password, ok = loginRequest.Data["password"].(string); !ok {
		err := fmt.Errorf("password is missing")
		log.Error("auth/login failed")

		fail.ReturnError(ctx, loginResponse, []string{err.Error()}, 400, log)
		return
	}

	for _, user := range Users {
		if user.Username == username && user.Password == password {
			token, err = crypto.UUID()
			if err != nil {
				err := fmt.Errorf("failed to generate UUID")
				log.Error("auth/login failed")

				fail.ReturnError(ctx, loginResponse, []string{err.Error()}, 400, log)
				return
			}
			role = string(user.Role)
		}
	}

	if token == "" {
		err := fmt.Errorf("login failed, username/password not found")
		log.Error("auth/login failed")

		fail.ReturnError(ctx, loginResponse, []string{err.Error()}, 400, log)
		return
	}

	Tokens = append(Tokens, model.Token{
		AccessToken: token,
		Role:        role,
	})

	ctx.Header(AccessTokenHeader, token)

	log.Info("login finished")

	loginResponse.Status = true
	ctx.JSON(200, loginResponse)
}

func init() {
	router.Router.Handle("POST", "auth/login", Login)
}
