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

func ReadHandler(ctx *gin.Context) {

	var (
		readRequest  = request.GetRequest{}
		readResponse = request.Response{}
	)

	// bind input data to request format
	err := ctx.ShouldBind(&readRequest)
	if err != nil {
		logger.Log.Error("Failed to bind input data",
			zap.Error(err),
		)

		fail.ReturnError(ctx, readResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	log := logger.Log.WithOptions(zap.Fields(
		zap.String("entity", readRequest.Entity),
		zap.Any("data", readRequest.Data),
	))

	log.Info("read started")

	// check authentication
	role, token := auth.GetIssuer(ctx)
	if role == model.Empty || token == "" {
		logger.Log.Error("failed to authorize user",
			zap.Error(err),
		)

		fail.ReturnError(ctx, readResponse, []string{"failed to authorize user"}, 400, logger.Log)
		return
	}

	for _, car := range Cars {
		if car.ID == readRequest.Data.ID {

			raw, err := json.Marshal(car)
			if err != nil {
				logger.Log.Error("json.Marshal failed",
					zap.Error(err),
				)

				fail.ReturnError(ctx, readResponse, []string{err.Error()}, 400, logger.Log)
				return
			}

			readResponse.Data = raw
			readResponse.Status = true

			break
		}
	}

	if !readResponse.Status {
		logger.Log.Error("failed to find car with given ID",
			zap.Error(err),
		)

		fail.ReturnError(ctx, readResponse, []string{"failed to find car with given ID"}, 400, logger.Log)
		return
	}

	log.Info("read finished")
	ctx.JSON(200, readResponse)
}

func ListHandler(ctx *gin.Context) {

	var (
		listRequest  = request.GetRequest{}
		listResponse = request.Response{}
		retrieveData = make([]model.Car, 0)
	)

	// bind input data to request format
	err := ctx.ShouldBind(&listRequest)
	if err != nil {
		logger.Log.Error("Failed to bind input data",
			zap.Error(err),
		)

		fail.ReturnError(ctx, listResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	log := logger.Log.WithOptions(zap.Fields(
		zap.String("entity", listRequest.Entity),
	))

	log.Info("list started")

	// check authentication
	role, token := auth.GetIssuer(ctx)
	if role == model.Empty || token == "" {
		logger.Log.Error("failed to authorize user",
			zap.Error(err),
		)

		fail.ReturnError(ctx, listResponse, []string{"failed to authorize user"}, 400, logger.Log)
		return
	}

	if listRequest.Metadata.Limit < 0 ||
		listRequest.Metadata.Limit > len(Cars) {
		logger.Log.Error("invalid limit values in metadata",
			zap.Int("limit", listRequest.Metadata.Limit),
		)

		fail.ReturnError(ctx, listResponse, []string{"invalid limit values in metadata"}, 400, logger.Log)
		return
	}

	if listRequest.Metadata.Offset < 0 ||
		listRequest.Metadata.Offset > len(Cars) {
		logger.Log.Error("invalid offset values in metadata",
			zap.Int("offset", listRequest.Metadata.Limit),
		)

		fail.ReturnError(ctx, listResponse, []string{"invalid offset values in metadata"}, 400, logger.Log)
		return
	}

	// default limit
	if listRequest.Metadata.Limit == 0 {
		listRequest.Metadata.Limit = len(Cars)
	}

	retrieveData = Cars[listRequest.Metadata.Offset:listRequest.Metadata.Limit]

	raw, err := json.Marshal(retrieveData)
	if err != nil {
		logger.Log.Error("json.Marshal failed",
			zap.Error(err),
		)

		fail.ReturnError(ctx, listResponse, []string{err.Error()}, 400, logger.Log)
		return
	}

	log.Info("list finished")

	listResponse.Data = raw
	listResponse.Total = len(Cars)
	listResponse.Status = true
	ctx.JSON(200, listResponse)
}

func init() {
	router.Router.Handle("POST", "/car/read", ReadHandler)
	router.Router.Handle("POST", "/car/list", ListHandler)
}
