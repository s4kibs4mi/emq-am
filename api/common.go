package api

import (
	"net/http"
	"github.com/spf13/viper"
	"encoding/json"
)

const (
	AppKey    = "app_key"
	AppSecret = "app_secret"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Details string      `json:"details,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ServeJSON(w http.ResponseWriter, result interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(result)
}

func ParseResponse(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func AppAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appKey := viper.GetString("security.key")
		appSecret := viper.GetString("security.secret")
		headerAppKey := r.Header.Get(AppKey)
		headerAppSecret := r.Header.Get(AppSecret)
		if appKey == headerAppKey && appSecret == headerAppSecret {
			h.ServeHTTP(w, r)
			return
		}
		ServeJSON(w, APIResponse{
			Code:    http.StatusUnauthorized,
			Details: "Authorization header missing.",
		}, http.StatusUnauthorized)
	}
}

func DefaultAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func MemberAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func AdminAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
