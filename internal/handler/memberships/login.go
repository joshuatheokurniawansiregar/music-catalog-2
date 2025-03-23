package memberships

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
)

func (h *Handler) Login(context *gin.Context){
	var requestLogin memberships.LoginRequest
	if err := context.ShouldBindJSON(&requestLogin); err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if(requestLogin == memberships.LoginRequest{}){
		context.JSON(http.StatusBadRequest, gin.H{
			"errror":errors.New("body request cannot be null or equal {}").Error(),
		})
		return
	}

	accessToken, err := h.serviceInterface.Login(&requestLogin)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusAccepted, memberships.LoginResponse{
		AccessToken: accessToken,
	})
}