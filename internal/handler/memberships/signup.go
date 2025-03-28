package memberships

import (
	"errors"
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

	if(request == memberships.SignUpRequest{}){
		context.JSON(http.StatusBadRequest, gin.H{
			"errror":errors.New("body request cannot be null or equal {}").Error(),
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