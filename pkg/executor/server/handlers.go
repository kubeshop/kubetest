package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

func (e *Executor) StartExecution() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request testkube.ExecutorStartRequest
		err := json.Unmarshal(c.Body(), &request)
		if err != nil {
			return e.Error(c, http.StatusBadRequest, err)
		}

		execution := testkube.ExecutorStartRequestToExecution(request)

		result := testkube.NewQueuedResult()
		execution.ExecutionResult = &result

		execution.WithContent(request.Content).
			WithParams(request.Params)

		if request.Repository != nil {
			execution.WithRepositoryData(
				request.Repository.Uri,
				request.Repository.Branch,
				request.Repository.Path,
			)
		}

		err = e.Repository.Insert(context.Background(), execution)
		if err != nil {
			return e.Error(c, http.StatusInternalServerError, err)

		}

		e.Log.Infow("starting new execution", "execution", execution)
		c.Response().Header.SetStatusCode(201)
		return c.JSON(execution)
	}
}

func (e *Executor) GetExecution() fiber.Handler {
	return func(c *fiber.Ctx) error {
		execution, err := e.Repository.Get(context.Background(), c.Params("id"))
		if err != nil {
			return e.Error(c, http.StatusInternalServerError, err)
		}

		return c.JSON(execution)
	}
}
