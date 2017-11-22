package api

import (
	"net/http"
	"github.com/s4kibs4mi/emq-am/data"
	"gopkg.in/mgo.v2/bson"
)

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	user := &data.User{}
	parseErr := ParseFromStringBody(r, user)
	if parseErr != nil || !bson.IsObjectIdHex(user.UserName) {
		ServeJSON(w, APIResponse{
			Code:   http.StatusBadRequest,
			Errors: parseErr,
		}, http.StatusBadRequest)
		return
	}
	session := data.Session{}
	session.UserId = bson.ObjectIdHex(user.UserName)
	session.AccessToken = user.Password
	if session.Find() {
		ServeJSON(w, APIResponse{
			Code: http.StatusOK,
		}, http.StatusOK)
		return
	}
	ServeJSON(w, APIResponse{
		Code: http.StatusUnauthorized,
	}, http.StatusUnauthorized)
}

func HasBroadcastPermission(w http.ResponseWriter, r *http.Request) {
	params := &data.ACLParams{}
	parseErr := ParseACLParams(r, params)
	if parseErr != nil || !bson.IsObjectIdHex(params.UserId) {
		ServeJSON(w, APIResponse{
			Code:   http.StatusBadRequest,
			Errors: parseErr,
		}, http.StatusBadRequest)
		return
	}
	user := &data.User{}
	user.Id = bson.ObjectIdHex(params.UserId)
	if params.Access == data.Publish && user.HasPublishPermission(params.Topic) {
		ServeJSON(w, APIResponse{
			Code: http.StatusOK,
		}, http.StatusOK)
		return
	} else if params.Access == data.Subscribe && user.HasSubscribePermission(params.Topic) {
		ServeJSON(w, APIResponse{
			Code: http.StatusOK,
		}, http.StatusOK)
		return
	}
	ServeJSON(w, APIResponse{
		Code: http.StatusUnauthorized,
	}, http.StatusUnauthorized)
}
