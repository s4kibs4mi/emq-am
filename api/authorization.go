package api

import (
	"net/http"
	"github.com/s4kibs4mi/emq-am/data"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/s4kibs4mi/emq-am/utils"
)

func HasBroadcastPermission(w http.ResponseWriter, r *http.Request) {
	params := &data.ACLParams{}
	parseErr := ParseACLParams(r, params)
	if parseErr != nil {
		ServeJSON(w, APIResponse{
			Code:   http.StatusBadRequest,
			Errors: parseErr,
		}, http.StatusBadRequest)
		return
	}
	fmt.Println(params)
	user := &data.User{}
	user.UserName = params.UserName
	if params.Access == data.MQTopicDirectionPublish && user.HasPublishPermission(params.Topic) {
		ServeJSON(w, APIResponse{
			Code: http.StatusOK,
		}, http.StatusOK)
		return
	} else if params.Access == data.MQTopicDirectionSubscribe && user.HasSubscribePermission(params.Topic) {
		ServeJSON(w, APIResponse{
			Code: http.StatusOK,
		}, http.StatusOK)
		return
	}
	ServeJSON(w, APIResponse{
		Code: http.StatusUnauthorized,
	}, http.StatusUnauthorized)
}

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
			Details: "Publish topics updated",
		}, http.StatusOK)
		return
	}
	ServeJSON(w, APIResponse{
		Code:    http.StatusInternalServerError,
		Details: "Couldn't update publish topics",
	}, http.StatusInternalServerError)
}
