package api

import (
	"net/http"
	"github.com/s4kibs4mi/emq-am/data"
	"github.com/s4kibs4mi/emq-am/utils"
	"gopkg.in/asaskevich/govalidator.v4"
	"github.com/spf13/viper"
	"github.com/satori/go.uuid"
	"time"
	"gopkg.in/mgo.v2/bson"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	isRegistrationEnabled := viper.GetBool("security.registration_enabled")
	if !isRegistrationEnabled {
		ServeJSON(w, APIResponse{
			Code:    http.StatusServiceUnavailable,
			Details: "User registration disabled",
		}, http.StatusServiceUnavailable)
		return
	}
	user := &data.User{}
	parseErr := ParseResponse(r, user)
	if parseErr != nil {
		ServeJSON(w, APIResponse{
			Code:    http.StatusBadRequest,
			Details: "Couldn't parse request body",
		}, http.StatusBadRequest)
		return
	}
	var validationErr []string
	userNameLen := len(user.UserName)
	if userNameLen < 3 || userNameLen > 1000 {
		validationErr = append(validationErr, "user_name must be between 3 to 1000")
	}
	passwordLen := len(user.Password)
	if passwordLen < 8 || passwordLen > 100 {
		validationErr = append(validationErr, "password must be between 8 to 100")
	}
	if !govalidator.IsEmail(user.Email) {
		validationErr = append(validationErr, "email must be a valid email address")
	}
	if len(validationErr) > 0 {
		ServeJSON(w, APIResponse{
			Code:   http.StatusUnprocessableEntity,
			Errors: validationErr,
		}, http.StatusUnprocessableEntity)
		return
	}
	if !user.IsEmailAvailable() {
		ServeJSON(w, APIResponse{
			Code:    http.StatusUnprocessableEntity,
			Details: "Email address already registered",
		}, http.StatusUnprocessableEntity)
		return
	}
	if !user.IsUserNameAvailable() {
		ServeJSON(w, APIResponse{
			Code:    http.StatusUnprocessableEntity,
			Details: "Username already taken",
		}, http.StatusUnprocessableEntity)
		return
	}
	if user.Count() == 0 {
		user.Type = data.UserTypeAdmin
	} else {
		user.Type = data.UserTypeDefault
	}
	user.Password = utils.MakePassword(user.Password)
	user.Id = bson.NewObjectId()
	if user.Password == "" || !user.Save() {
		ServeJSON(w, APIResponse{
			Code:    http.StatusInternalServerError,
			Details: "Couldn't store data",
		}, http.StatusInternalServerError)
		return
	}
	user.Password = ""
	ServeJSON(w, APIResponse{
		Code:    http.StatusOK,
		Details: "User successfully created",
		Data:    user,
	}, http.StatusOK)
	return
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	user := &data.User{}
	parseErr := ParseFromStringBody(r, user)
	if parseErr != nil {
		ServeJSON(w, APIResponse{
			Code:   http.StatusBadRequest,
			Errors: parseErr,
		}, http.StatusBadRequest)
		return
	}
	if user.HasValidCredentials() {
		ServeJSON(w, APIResponse{
			Code: http.StatusOK,
		}, http.StatusOK)
		return
	}
	ServeJSON(w, APIResponse{
		Code: http.StatusUnauthorized,
	}, http.StatusUnauthorized)
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
	user := &data.User{}
	parseErr := ParseResponse(r, user)
	if parseErr != nil {
		ServeJSON(w, APIResponse{
			Code:   http.StatusBadRequest,
			Errors: parseErr,
		}, http.StatusBadRequest)
		return
	}
	if user.HasValidCredentials() {
		session := &data.Session{
			Id:           bson.NewObjectId(),
			UserId:       user.Id,
			AccessToken:  uuid.NewV4().String(),
			RefreshToken: uuid.NewV4().String(),
			CreatedAt:    time.Now(),
			ExpireAt:     time.Now().Add(time.Hour * 24),
		}
		if session.Save() {
			ServeJSON(w, APIResponse{
				Code: http.StatusOK,
				Data: session,
			}, http.StatusOK)
			return
		}
		ServeJSON(w, APIResponse{
			Code:    http.StatusInternalServerError,
			Details: "Couldn't generate session",
		}, http.StatusInternalServerError)
		return
	}
	ServeJSON(w, APIResponse{
		Code: http.StatusUnauthorized,
	}, http.StatusUnauthorized)
}
