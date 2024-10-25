// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/v1/health)
	GetApiV1Health(c *gin.Context)

	// (GET /api/v1/messages/{location_id})
	GetApiV1MessagesLocationId(c *gin.Context, locationId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetApiV1Health operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Health(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiV1Health(c)
}

// GetApiV1MessagesLocationId operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1MessagesLocationId(c *gin.Context) {

	var err error

	// ------------- Path parameter "location_id" -------------
	var locationId string

	err = runtime.BindStyledParameter("simple", false, "location_id", c.Param("location_id"), &locationId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter location_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiV1MessagesLocationId(c, locationId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/api/v1/health", wrapper.GetApiV1Health)
	router.GET(options.BaseURL+"/api/v1/messages/:location_id", wrapper.GetApiV1MessagesLocationId)
}
