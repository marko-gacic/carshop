package crud

import (
	"carshop/internal/model"
	"carshop/internal/request"
	"carshop/internal/route/auth"
	"carshop/internal/route/fail"
	"carshop/internal/server/router"
	"carshop/pkg/crypto"
	"carshop/pkg/logger"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateHandler(ctx *gin.Context) {
	var (
		createRequest  = request.Request{}
		createResponse = request.Response{}
		car            = model.Car{}
	)

	// bind input data to request format
	err := ctx.ShouldBind(&createRequest)
	if err != nil {
		logger.Log.Error("failed to bind input data",
			zap.Error(err),
		)

		fail.ReturnError(ctx, createResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	log := logger.Log.WithOptions(zap.Fields(
		zap.String("entity", createRequest.Entity),
		zap.Any("data", createRequest.Data),
	))

	log.Info("create started")

	// check authentication
	role, token := auth.GetIssuer(ctx)
	if role == model.Empty || token == "" {
		logger.Log.Error("failed to authorize user",
			zap.Error(err),
		)

		fail.ReturnError(ctx, createResponse, []string{"failed to authorize user"}, 400, logger.Log)
		return
	}

	raw, err := json.Marshal(createRequest.Data)
	if err != nil {
		logger.Log.Error("json.Marshal failed",
			zap.Error(err),
		)

		fail.ReturnError(ctx, createResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	err = json.Unmarshal(raw, &car)
	if err != nil {
		logger.Log.Error("json.Unmarshal failed",
			zap.Error(err),
		)

		fail.ReturnError(ctx, createResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	if car.ID == "" {
		car.ID, err = crypto.UUID()
		if err != nil {
			logger.Log.Error("system error",
				zap.Error(err),
			)

			fail.ReturnError(ctx, createResponse, []string{err.Error()}, 400, logger.Log)
			return
		}
	}

	car.Active = true
	for _, nextCar := range Cars {
		if nextCar.ID == car.ID {
			log.Warn("an entity with the given ID already exists",
				zap.String("id", car.ID),
			)

			fail.ReturnError(ctx, createResponse, []string{"an entity with the given ID already exists"}, 400, logger.Log)
			return
		}
	}

	Cars = append(Cars, car)

	log.Info("create finished")

	//createResponse.Data = response
	createResponse.Status = true
	ctx.JSON(200, createResponse)
}

func init() {
	router.Router.Handle("POST", "/car/create", CreateHandler)
}
