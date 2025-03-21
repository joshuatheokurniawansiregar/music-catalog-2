package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
)

func (h *Handler) Signup(context *gin.Context) {
	var request memberships.SignUpRequest
	err:= context.ShouldBindJSON(&request)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"errror":err.Error(),
		})
		return
	}

	err = h.serviceInterface.Signup(request)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.Status(http.StatusCreated)
}