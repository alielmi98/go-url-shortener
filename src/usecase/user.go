package usecase

import (
	"log"

	"github.com/alielmi98/go-url-shortener/api/dto"
	"github.com/alielmi98/go-url-shortener/config"
	"github.com/alielmi98/go-url-shortener/constants"
	"github.com/alielmi98/go-url-shortener/data/models"
	"github.com/alielmi98/go-url-shortener/data/repository"
	"github.com/alielmi98/go-url-shortener/pkg/service_errors"
	"github.com/alielmi98/go-url-shortener/services"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	cfg     *config.Config
	repo    *repository.UserRepository
	service *services.TokenService
}

func NewUserUsecase(cfg *config.Config) *UserUsecase {
	return &UserUsecase{
		cfg:     cfg,
		repo:    repository.NewUserRepository(),
		service: services.NewTokenService(cfg),
	}
}

// Register by username
func (s *UserUsecase) RegisterByUsername(req *dto.RegisterUserByUsernameRequest) error {
	u := models.User{Username: req.Username, Email: req.Email}

	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Caller:%s Level:%s Msg:%s", constants.General, constants.HashPassword, err.Error())
		return err
	}
	u.Password = string(hp)

	err = s.repo.Create(&u)
	if err != nil {
		return err
	}
	return nil

}

func (s *UserUsecase) LoginByUsername(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	u, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UsernameOrPasswordInvalid}
	}

	token := &dto.TokenDto{UserId: u.Id, Username: u.Username, Email: u.Email}
	return s.service.GenerateToken(token)
}
