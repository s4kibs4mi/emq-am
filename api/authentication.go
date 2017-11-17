package api

import (
	"net/http"
	"github.com/s4kibs4mi/emq-am/data"
	"github.com/s4kibs4mi/emq-am/utils"
	"gopkg.in/asaskevich/govalidator.v4"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
