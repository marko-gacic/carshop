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

func DeleteHandler(ctx *gin.Context) {
	var (
		deleteRequest  = request.Request{}
		deleteResponse = request.Response{}
		car            = model.Car{}
		found          bool
	)

	// bind input data to request format
	err := ctx.ShouldBind(&deleteRequest)
	if err != nil {
		logger.Log.Error("failed to bind input data",
			zap.Error(err),
		)

		fail.ReturnError(ctx, deleteResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	log := logger.Log.WithOptions(zap.Fields(
		zap.String("entity", deleteRequest.Entity),
		zap.Any("data", deleteRequest.Data),
	))

	log.Info("delete started")

	// check authentication
	role, token := auth.GetIssuer(ctx)
	if role == model.Empty || token == "" {
		logger.Log.Error("failed to authorize user",
			zap.Error(err),
		)

		fail.ReturnError(ctx, deleteResponse, []string{"failed to authorize user"}, 400, logger.Log)
		return
	}

	raw, err := json.Marshal(deleteRequest.Data)
	if err != nil {
		logger.Log.Error("json.Marshal failed",
			zap.Error(err),
		)

		fail.ReturnError(ctx, deleteResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	err = json.Unmarshal(raw, &car)
	if err != nil {
		logger.Log.Error("json.Unmarshal failed",
			zap.Error(err),
		)

		fail.ReturnError(ctx, deleteResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	tempCars := make([]model.Car, 0)
	for _, nextCar := range Cars {
		if nextCar.ID == car.ID {
			found = true
			continue
		}
		tempCars = append(tempCars, nextCar)
	}
	Cars = tempCars

	if !found {
		log.Warn("nothing was deleted for the given id",
			zap.String("id", car.ID),
		)

		fail.ReturnError(ctx, deleteResponse, []string{"nothing was deleted for the given id"}, 400, logger.Log)
		return
	}

	log.Info("delete finished")

	//deleteResponse.Data = response
	deleteResponse.Status = true
	ctx.JSON(200, deleteResponse)
}

func init() {
	router.Router.Handle("POST", "/car/delete", DeleteHandler)

}
