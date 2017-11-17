package api

import (
	"net/http"
	"github.com/spf13/viper"
	"encoding/json"
	"github.com/s4kibs4mi/emq-am/data"
)

const (
	AppKey    = "app_key"
	AppSecret = "app_secret"
)

func ServeJSON(w http.ResponseWriter, result interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(result)
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
		ServeJSON(w, data.User{}, http.StatusUnauthorized)
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
