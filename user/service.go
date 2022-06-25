package user

import (
	"financial-tracker-be/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Service interface {
	UserRegister(userBody UserRegisterBody) (User, error)
	UserLoginWithJWT(userAuth UserLoginRequest) (string, error)
	UserLoginAes(userAuth UserLoginRequest) (string, error)
}

type service struct {
	repository Repository
}

func UserService(repository Repository) *service {
	return &service{repository}
}

func (s service) UserRegister(userBody UserRegisterBody) (User, error) {

	digestUser, err := utils.EncryptStringToBase64(userBody.Email+userBody.Password, "alpha")

	if err != nil {
		return User{}, err
	}

	user := User{
		ID:                       uuid.Must(uuid.NewRandom()),
		Name:                     userBody.Name,
		Email:                    userBody.Email,
		DigestUserAuth:           digestUser,
		ConfirmationToken:        utils.GenerateRandomChars(32),
		ConfirmationTokenExpired: time.Now().Add(time.Hour * 24),
	}

	newUser, err := s.repository.CreateUser(user)

	return newUser, err
}

func (s service) UserLoginWithJWT(userAuth UserLoginRequest) (string, error) {
	var hmacSampleSecret []byte
	var userDigest = utils.DigestStringUsingMD5(userAuth.Email + userAuth.Password)

	user, err := s.repository.VerifyUser(userDigest)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     userAuth.Email,
		"name":      user.Name,
		"timestamp": time.Now(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

func (s service) UserLoginAes(userAuth UserLoginRequest) (string, error) {
	digestUser, err := utils.EncryptStringToBase64(userAuth.Email+userAuth.Password, "alpha")

	if err != nil {
		return "error to encrypt auth", err
	}

	user, err := s.repository.VerifyUser(digestUser)

	if err != nil {
		return "error to verify user", err
	}

	randomStrings := utils.GenerateRandomStrings(15)
	currentEpochTime := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	maskedDigest := randomStrings + user.DigestUserAuth + currentEpochTime

	fmt.Println(maskedDigest)
	tokenString, err := utils.EncryptStringToBase64(maskedDigest, "beta")

	return tokenString, err
}
