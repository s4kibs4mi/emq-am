package api

import (
	"net/http"
	"github.com/s4kibs4mi/emq-am/data"
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
