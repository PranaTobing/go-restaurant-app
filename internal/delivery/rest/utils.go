package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/rocksus/go-restaurant-app/internal/model"
)

func GetSessionData(r *http.Request) (model.UserSession, error) {
	authString := r.Header.Get("Authorization")
	splitString := strings.Split(authString, " ")
	if len(splitString) != 2 {
		return model.UserSession{}, errors.New("unauthorized")
	}
	accessString := splitString[1]

	return model.UserSession{
		JWTToken: accessString,
	}, nil
}
