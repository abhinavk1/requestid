# Request ID Gin middleware 

Request ID middleware for Gin Framework. 
Adds an identifier to the response using the `X-Request-ID` header. 
Passes the `X-Request-ID` value back to the caller if it's sent in the request headers.

## Config

define your custom generator function:

```go
func main() {

	r := gin.New()

	r.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return "test"
		},
	}))

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```

## Example

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abhinavk1/requestid"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	r.Use(requestid.New())

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```

How to get the request identifier:

```go
// Example / request.
r.GET("/", func(c *gin.Context) {
	c.String(http.StatusOK, "id:"+requestid.Get(c))
})
```

## Using a Custom Request Id Header

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abhinavk1/requestid"
	"github.com/gin-gonic/gin"
)

const customRequestIdHeader = "X-Request-Id-Custom" 

func main() {

	r := gin.New()

    r.Use(requestid.New(requestid.Config{
        CustomRequestIdHeader: customRequestIdHeader,
    }))

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```