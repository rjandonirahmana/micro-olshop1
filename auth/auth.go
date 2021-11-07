package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(customer_id uint) (string, error)
	ValidateToken(token string) (uint, error)
	GenerateTokenSeller(seller_id uint) (string, error)
	ValidateTokenSeller(encodedToken string) (uint, error)
}

type jwtService struct {
	SECRET_KEY  []byte
	SECRET_KEY2 []byte
}

func NewService(secret_key_customer, secret_key_seller string) *jwtService {
	return &jwtService{SECRET_KEY: []byte(secret_key_customer), SECRET_KEY2: []byte(secret_key_seller)}
}

func (j *jwtService) GenerateToken(customer_id uint) (string, error) {

	//claim adalah payload data jwt
	claim := jwt.MapClaims{}
	claim["customer_id"] = customer_id
	claim["exp"] = time.Now().Add(time.Hour * 8).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(j.SECRET_KEY)

	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (j *jwtService) ValidateToken(encodedToken string) (uint, error) {
	tokenclaim, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("INVALID ERROR")
		}

		return []byte(j.SECRET_KEY), nil

	})
	if err != nil {
		return 0, err
	}
	claim, ok := tokenclaim.Claims.(jwt.MapClaims)
	if !ok || !tokenclaim.Valid {
		return 0, errors.New("unauthorized")
	}

	id := uint(claim["customer_id"].(float64))

	return id, nil
}

func (j *jwtService) GenerateTokenSeller(seller_id uint) (string, error) {

	//claim adalah payload data jwt
	claim := jwt.MapClaims{}
	claim["seller_id"] = seller_id
	claim["exp"] = time.Now().Add(time.Hour * 8).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(j.SECRET_KEY2)

	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (j *jwtService) ValidateTokenSeller(encodedToken string) (uint, error) {
	tokenclaim, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("INVALID ERROR")
		}

		return []byte(j.SECRET_KEY2), nil

	})
	if err != nil {
		return 0, err
	}

	claim, ok := tokenclaim.Claims.(jwt.MapClaims)
	if !ok || !tokenclaim.Valid {
		return 0, err
	}

	id := uint(claim["seller_id"].(float64))

	return id, nil
}