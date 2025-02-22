package services

import (
	"time"

	"github.com/alielmi98/go-url-shortener/api/dto"
	"github.com/alielmi98/go-url-shortener/config"
	"github.com/alielmi98/go-url-shortener/constants"
	"github.com/alielmi98/go-url-shortener/pkg/service_errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	cfg *config.Config
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

func (s *TokenService) GenerateToken(token *dto.TokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}
	td.AccessTokenExpireTime = time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(s.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	atc := jwt.MapClaims{}

	atc[constants.UserIdKey] = token.UserId
	atc[constants.UsernameKey] = token.Username
	atc[constants.EmailKey] = token.Email
	atc[constants.ExpireTimeKey] = td.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	var err error
	td.AccessToken, err = at.SignedString([]byte(s.cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}

	rtc[constants.UserIdKey] = token.UserId
	rtc[constants.UsernameKey] = token.Username
	rtc[constants.EmailKey] = token.Email
	rtc[constants.ExpireTimeKey] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	td.RefreshToken, err = rt.SignedString([]byte(s.cfg.JWT.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
		}
		return []byte(s.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	// Check if the token is valid and not expired
	if claims, ok := at.Claims.(jwt.MapClaims); ok && at.Valid {
		// Check the expiration time
		if float64(time.Now().Unix()) > claims[constants.ExpireTimeKey].(float64) {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
		}
		return at, nil
	}

	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}

func (s *TokenService) RefreshToken(c *gin.Context) (*dto.TokenDetail, error) {
	refreshToken, err := c.Cookie(constants.RefreshTokenCookieName)
	if err != nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.InvalidRefreshToken}
	}

	claims, err := s.GetClaims(refreshToken)
	if err != nil {
		return nil, err
	}

	tokenDto := &dto.TokenDto{
		UserId:   int(claims[constants.UserIdKey].(float64)),
		Username: claims[constants.UsernameKey].(string),
		Email:    claims[constants.EmailKey].(string),
	}
	newTokenDetail, err := s.GenerateToken(tokenDto)
	if err != nil {
		return nil, err
	}

	c.SetCookie(constants.RefreshTokenCookieName, newTokenDetail.RefreshToken, int(s.cfg.JWT.RefreshTokenExpireDuration*60), "/", s.cfg.Server.Domin, true, true)

	return newTokenDetail, nil
}
