package utils

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ishanshre/Go-blog/api/v1/models"
)

func GenerateTokens(id int) (*models.LoginResponse, error) {
	/*
		Generating access and refresh tokens
	*/
	access_claims := jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(time.Minute * 15)).Unix(),
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
		"user_id":   id,
	}
	secret := os.Getenv("JWT_SECRET")
	ss := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)
	access_token, err := ss.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	refresh_claims := jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(time.Hour * 1)).Unix(),
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
		"user_id":   id,
	}
	rs := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)
	refresh_token, err := rs.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &models.LoginResponse{
		ID:           id,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}, nil
}

func ExtractToken(r *http.Request) (string, error) {
	/*
		This method extracts token from the header
	*/
	bearerToken := r.Header.Get("Authorization")
	tokenString := strings.Split(bearerToken, " ")
	if len(tokenString) == 2 {
		return tokenString[1], nil
	}
	return "", fmt.Errorf("invalid token")
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	/*
		This method verifies if the token is expired or not
	*/
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := ExtractToken(r)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) (jwt.MapClaims, error) {
	/*
		Check if the token is valid
	*/
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("token not valid")
	}
	return claims, nil
}

func ExractTokenMetaData(r *http.Request) (*models.TokenMetaData, error) {
	/*
		Extract data from the token
	*/
	claims, err := TokenValid(r)
	if err != nil {
		return nil, err
	}
	userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 0)
	if err != nil {
		return nil, err
	}
	return &models.TokenMetaData{ID: int(userId)}, nil
}

func VerifyUser(id int, r *http.Request) error {
	/*
		returns true if id match with id extracted from the token
	*/
	claims, err := TokenValid(r)
	if err != nil {
		return err
	}
	if int64(id) != int64(claims["user_id"].(float64)) {
		return fmt.Errorf("permission denied")
	}
	return nil
}
