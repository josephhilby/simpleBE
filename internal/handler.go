// Gin handlers, processes incoming HTTP requests

package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	v1 "simpleBE/api/pb/v1"
)

// Handler struct holds a reference to the service layer
type Handler struct {
	svc *Service
}

// NewHandler constructs a Handler with a given service instance
func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// GetHello handles GET /api/hello requests
func (h *Handler) GetHello(c *gin.Context) {
	// Call the service to fetch the hello message
	msg, err := h.svc.GetHelloMessage()
	if err != nil {
		// If there's an error, respond with a 500 and the error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Construct the protobuf response object
	resp := &v1.HelloReply{Message: msg}

	// Marshal the response to JSON using protobuf-safe marshaler
	jsonBytes, _ := protojson.Marshal(resp)

	// Send the JSON response with a 200 status code
	c.Data(http.StatusOK, "application/json", jsonBytes)
}
