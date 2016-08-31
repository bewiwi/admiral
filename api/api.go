package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(address string, port int) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/version", getVersion)

	v1 := r.Group("/v1")
	{
		// Registry events notification endpoint
		v1.GET("/events", getEvents)
		v1.POST("/events", postEvents)

		// Namespace endpoints
		v1.GET("/namespaces", getNamespaces)
		v1.GET("/namespace/:id", getNamespace)
		v1.POST("/namespace", postNamespace)
		v1.DELETE("/namespace/:id", deleteNamespace)
		v1.PATCH("/namespace/:id", patchNamespace)
	}

	r.Run(fmt.Sprintf("%s:%d", address, port))
}
