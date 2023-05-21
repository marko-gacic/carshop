package crud

import (
	"carshop/internal/model"
	"carshop/internal/request"
	"carshop/internal/route/auth"
	"carshop/internal/route/fail"
	"carshop/internal/server/router"
	"carshop/pkg/logger"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UpdateHandler(ctx *gin.Context) {
	var (
		updateRequest  = request.Request{}
		updateResponse = request.Response{}
		car            = model.Car{}
		found          bool
	)

	// bind input data to request format
	err := ctx.ShouldBind(&updateRequest)
	if err != nil {
		logger.Log.Error("failed to bind input data",
			zap.Error(err),
		)

		fail.ReturnError(ctx, updateResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	log := logger.Log.WithOptions(zap.Fields(
		zap.String("entity", updateRequest.Entity),
		zap.Any("data", updateRequest.Data),
	))

	log.Info("update started")

	// check authentication
	role, token := auth.GetIssuer(ctx)
	if role == model.Empty || token == "" {
		logger.Log.Error("failed to authorize user",
			zap.Error(err),
		)

		fail.ReturnError(ctx, updateResponse, []string{"failed to authorize user"}, 400, logger.Log)
		return
	}

	raw, err := json.Marshal(updateRequest.Data)
	if err != nil {
		logger.Log.Error("json.Marshal failed",
			zap.Error(err),
		)

		fail.ReturnError(ctx, updateResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	err = json.Unmarshal(raw, &car)
	if err != nil {
		logger.Log.Error("json.Unmarshal failed",
			zap.Error(err),
		)

		fail.ReturnError(ctx, updateResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	// Find the car in the existing list and update its details
	for i, existingCar := range Cars {
		if existingCar.ID == car.ID {
			Cars[i] = car
			found = true
			break
		}
	}

	if !found {
		log.Warn("no car found for the given id",
			zap.String("id", car.ID),
		)

		fail.ReturnError(ctx, updateResponse, []string{"no car found for the given id"}, 400, logger.Log)
		return
	}

	log.Info("update finished")

	updateResponse.Status = true
	ctx.JSON(200, updateResponse)
}

func init() {
	router.Router.Handle("POST", "/car/update", UpdateHandler)
}
