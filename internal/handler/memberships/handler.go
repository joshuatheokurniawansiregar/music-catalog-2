package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships
type serviceInterface interface {
	Signup(request memberships.SignUpRequest)error
}

type Handler struct {
	*gin.Engine
	serviceInterface serviceInterface
}

func NewHandler(engine *gin.Engine, serviceInterface serviceInterface)*Handler{
	return &Handler{
		Engine: engine,
		serviceInterface: serviceInterface,
	}
}

func(h *Handler) RegisterRoute(){
	var route *gin.RouterGroup = h.Group("/api/v1/memberships")
	route.POST("/sign_up", h.Signup)
}
