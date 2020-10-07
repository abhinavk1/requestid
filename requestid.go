package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	headerXRequestID = "X-Request-ID"
	requestHeaderKey = "requestHeaderKey"
)

// Config defines the config for RequestID middleware
type Config struct {
	// Generator defines a function to generate an ID.
	// Optional. Default: func() string {
	//   return uuid.New().String()
	// }
	Generator func() string

	CustomRequestIdHeader string
}

// New initializes the RequestID middleware.
func New(config ...Config) gin.HandlerFunc {

	var cfg Config
	var requestIdHeader string

	if len(config) > 0 {
		cfg = config[0]
	}

	// Set config default values
	if cfg.Generator == nil {
		cfg.Generator = func() string {
			return uuid.New().String()
		}
	}

	if cfg.CustomRequestIdHeader != "" {
		requestIdHeader = cfg.CustomRequestIdHeader
	} else {
		requestIdHeader = headerXRequestID
	}

	return func(ctx *gin.Context) {
		ctx.Set(requestHeaderKey, requestIdHeader)

		// Get id from request
		rid := ctx.GetHeader(requestIdHeader)
		if rid == "" {
			rid = cfg.Generator()
		}

		// Set the id to ensure that the request Id is in the response
		ctx.Header(requestIdHeader, rid)
		ctx.Next()
	}
}

// Get returns the request identifier
func Get(ctx *gin.Context) string {

	requestIdHeader := ctx.GetString(requestHeaderKey)
	return ctx.Writer.Header().Get(requestIdHeader)
}
