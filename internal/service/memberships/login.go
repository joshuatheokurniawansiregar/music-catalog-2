package memberships

import (
	"errors"

	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func(s *Service) Login(request *memberships.LoginRequest)(string, error){
	user, err := s.repositoryInterface.GetUser(request.Email, "", 0)
	if err != nil && err != gorm.ErrRecordNotFound{
		return "", errors.New("error get user from database")
	}

	if user == nil{
		return "", errors.New("email does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil{
		return "", err
	}

	accessToken, err := jwt.CreateToken(int64(user.ID), user.Username, s.cfg.Service.SecretJWT)
	if err != nil{
		return "", errors.New("failed to get jwt token")
	}
	return accessToken, nil
}