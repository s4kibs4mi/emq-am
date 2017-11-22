package api

import (
	"net/http"
	"strconv"
	"github.com/s4kibs4mi/emq-am/data"
	"gopkg.in/mgo.v2/bson"
	"github.com/s4kibs4mi/emq-am/utils"
)

/**
 * := Coded with love by Sakib Sami on 20/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func CreatePublishTopic(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get(UserId)
	params := &data.ACLParams{}
	parseErr := ParseResponse(r, params)
	if parseErr != nil || params.Topic == "" {
		ServeJSON(w, APIResponse{
			Code:   http.StatusBadRequest,
			Errors: parseErr,
		}, http.StatusBadRequest)
		return
	}
	user := data.User{}
	user.Id = bson.ObjectIdHex(userId)
	if !user.FindById() {
		ServeJSON(w, APIResponse{
			Code:    http.StatusNotFound,
			Details: "User not found",
		}, http.StatusNotFound)
		return
	}
	if utils.IsItemExists(user.PublishTopics, params.Topic) {
		ServeJSON(w, APIResponse{
			Code:    http.StatusUnprocessableEntity,
			Details: "Topic already exists",
		}, http.StatusUnprocessableEntity)
		return
	}
	if user.AppendPublishPermission(params.Topic) {
		ServeJSON(w, APIResponse{
			Code:    http.StatusOK,
			Data:    user,
			Details: "Publish topic updated",
		}, http.StatusOK)
		return
	}
	ServeJSON(w, APIResponse{
		Code:    http.StatusInternalServerError,
		Details: "Couldn't update publish topic",
	}, http.StatusInternalServerError)
}

func RemovePublishTopic(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get(UserId)
	params := &data.ACLParams{}
	parseErr := ParseResponse(r, params)
	if parseErr != nil || params.Topic == "" {
		ServeJSON(w, APIResponse{
			Code:   http.StatusBadRequest,
			Errors: parseErr,
		}, http.StatusBadRequest)
		return
	}
	user := data.User{}
	user.Id = bson.ObjectIdHex(userId)
	if !user.FindById() {
		ServeJSON(w, APIResponse{
			Code:    http.StatusNotFound,
			Details: "User not found",
		}, http.StatusNotFound)
		return
	}
	if !utils.IsItemExists(user.PublishTopics, params.Topic) {
		ServeJSON(w, APIResponse{
			Code:    http.StatusNotFound,
			Details: "Topic doesn't exists",
		}, http.StatusNotFound)
		return
	}
	if user.DiscardPublishPermission(params.Topic) {
		ServeJSON(w, APIResponse{
			Code:    http.StatusOK,
			Data:    user,
			Details: "Publish topic removed",
		}, http.StatusOK)
		return
	}
	ServeJSON(w, APIResponse{
		Code:    http.StatusInternalServerError,
		Details: "Couldn't update publish topic",
	}, http.StatusInternalServerError)
}

func CreateSubscribeTopic(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get(UserId)
	params := &data.ACLParams{}
	parseErr := ParseResponse(r, params)
	if parseErr != nil || params.Topic == "" {
		ServeJSON(w, APIResponse{
			Code:   http.StatusBadRequest,
			Errors: parseErr,
		}, http.StatusBadRequest)
		return
	}
	user := data.User{}
	user.Id = bson.ObjectIdHex(userId)
	if !user.FindById() {
		ServeJSON(w, APIResponse{
			Code:    http.StatusNotFound,
			Details: "User not found",
		}, http.StatusNotFound)
		return
	}
	if utils.IsItemExists(user.SubscribeTopics, params.Topic) {
		ServeJSON(w, APIResponse{
			Code:    http.StatusUnprocessableEntity,
			Details: "Topic already exists",
		}, http.StatusUnprocessableEntity)
		return
	}
	if user.AppendSubscribePermission(params.Topic) {
		ServeJSON(w, APIResponse{
			Code:    http.StatusOK,
			Data:    user,
			Details: "Subscribe topic updated",
		}, http.StatusOK)
		return
	}
	ServeJSON(w, APIResponse{
		Code:    http.StatusInternalServerError,
		Details: "Couldn't update subscribe topic",
	}, http.StatusInternalServerError)
}

func RemoveSubscribeTopic(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get(UserId)
	params := &data.ACLParams{}
	parseErr := ParseResponse(r, params)
	if parseErr != nil || params.Topic == "" {
		ServeJSON(w, APIResponse{
			Code:   http.StatusBadRequest,
			Errors: parseErr,
		}, http.StatusBadRequest)
		return
	}
	user := data.User{}
	user.Id = bson.ObjectIdHex(userId)
	if !user.FindById() {
		ServeJSON(w, APIResponse{
			Code:    http.StatusNotFound,
			Details: "User not found",
		}, http.StatusNotFound)
		return
	}
	if !utils.IsItemExists(user.SubscribeTopics, params.Topic) {
		ServeJSON(w, APIResponse{
			Code:    http.StatusNotFound,
			Details: "Topic doesn't exists",
		}, http.StatusNotFound)
		return
	}
	if user.DiscardSubscribePermission(params.Topic) {
		ServeJSON(w, APIResponse{
			Code:    http.StatusOK,
			Data:    user,
			Details: "Subscribe topic removed",
		}, http.StatusOK)
		return
	}
	ServeJSON(w, APIResponse{
		Code:    http.StatusInternalServerError,
		Details: "Couldn't update publish topic",
	}, http.StatusInternalServerError)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	currentPage := 0
	page := r.URL.Query().Get("page")
	n, err := strconv.Atoi(page)
	if err == nil {
		currentPage = n
	}
	user := data.User{}
	users := user.GetUserList(currentPage)
	ServeJSON(w, APIResponse{
		Code: http.StatusOK,
		Data: users,
	}, http.StatusOK)
}
