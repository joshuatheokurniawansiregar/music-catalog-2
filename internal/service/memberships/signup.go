package memberships

import (
	"errors"
	golog "log"

	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *Service) Signup(request memberships.SignUpRequest) error {
    existingUser, err := s.repositoryInterface.GetUser(request.Email, request.Username, 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error: error get user from database")
		return err
	}

	if existingUser != nil{
		golog.Printf("error: error when create user. email %s and username %s exist", request.Email, request.Username)
		return errors.New("emain and username exist")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil{
		log.Error().Err(err).Msg("error: error when hashing password")
		return err
	}
	var user memberships.User = memberships.User{
		Email: request.Email,
		Password: string(password),
		Username: request.Username,
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}

	err = s.repositoryInterface.CreateUser(user)
	if err != nil{
		log.Error().Err(err).Msg("error: error when create user")
		return err
	}
	return nil
}