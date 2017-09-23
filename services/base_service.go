package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/albimcleod/go-modish/authentication"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
)

// BaseService is an service to handle data requests
type BaseService struct {
	Name string
	//vapiRouter *VapiRouter
	//Session *repositories.DatabaseSession
}

// HandleError returns the error response
func (service *BaseService) HandleError(w http.ResponseWriter, err error, status int) bool {
	if err != nil {
		fmt.Printf("HandleError - %v\n", err)

		service.WriteHeaderStatus(w, status)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("HandleError.Encode - %v\n", err)
			panic(err)
		}
		return true
	}
	return false
}

// WriteHeaderStatus  writes the json content header
func (service *BaseService) WriteHeaderStatus(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
}

// HandleNotFound returns the not found status code
func (service *BaseService) HandleNotFound(w http.ResponseWriter) {
	service.WriteHeaderStatus(w, http.StatusNotFound)
}

// GetParam returns the value of a parameter
func (service *BaseService) GetParam(param string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[param]
}

// GetQuery returns the value of a query string
func (service *BaseService) GetQuery(param string, r *http.Request) string {
	vals := r.URL.Query()
	values, ok := vals[param] // Note type, not ID. ID wasn't specified anywhere.
	if ok {
		if len(values) > 0 {
			return values[0]
		}
	}
	return ""
}

// GetAuthorizationToken returns the authoriztion token
func (service *BaseService) GetAuthorizationToken(secret string, r *http.Request) (*jwt.Token, error) {
	var token string

	tokens, ok := r.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
	}

	// If the token is empty...
	if token == "" {
		return nil, fmt.Errorf("No Token")
	}

	return authentication.ParseToken(token, secret)
}
