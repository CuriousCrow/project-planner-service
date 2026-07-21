package wrappers

import (
	"context"
	"errors"
	"net/http"

	"github.com/CuriousCrow/project-planner-service/internal/metrics"
	"github.com/CuriousCrow/project-planner-service/typed_error"
	"github.com/gin-gonic/gin"
)

func handleTypedError(c *gin.Context, te typed_error.TypedError) {
	switch te.Type {
	case typed_error.NotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": te.Error()})
	case typed_error.BadRequest:
		c.JSON(http.StatusBadRequest, gin.H{"error": te.Error()})
	case typed_error.ServerError:
		c.JSON(http.StatusInternalServerError, gin.H{"error": te.Error()})
	}
}

func GetHandlerWrapper[Req any, Resp any](hFunc func(ctx context.Context, params gin.Params, req Req) (*Resp, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		metrics.RequestCounter.Inc()

		ctx := c.Request.Context()
		var req Req

		err := c.ShouldBindQuery(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := hFunc(ctx, c.Params, req)
		if err != nil {
			var sampleTypedErr = typed_error.TypedError{}

			if errors.As(err, &sampleTypedErr) {
				handleTypedError(c, sampleTypedErr)
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

func HandlerWrapper[Req any, Resp any](hFunc func(ctx context.Context, params gin.Params, req Req) (*Resp, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		metrics.RequestCounter.Inc()

		ctx := c.Request.Context()
		var req Req

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		resp, err := hFunc(ctx, c.Params, req)
		if err != nil {
			var sampleTypedErr = typed_error.TypedError{}

			if errors.As(err, &sampleTypedErr) {
				handleTypedError(c, sampleTypedErr)
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}
