package cmd

import (
	"github.com/spf13/cobra"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/spf13/viper"
	"github.com/s4kibs4mi/emq-am/api"
)

var ServeCmd = cobra.Command{
	Use:   "serve",
	Run:   ServeCmdExecute,
	Short: "Start server",
}

func ServeCmdExecute(command *cobra.Command, args []string) {
	router := mux.NewRouter()
	v1 := router.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/users", api.AppAuth(api.CreateUser)).Methods("POST")
	v1.HandleFunc("/users", api.AdminAuth(api.ListUsers)).Methods("GET")
	v1.HandleFunc("/auth", api.CheckLogin).Methods("POST")
	v1.HandleFunc("/acl", api.HasBroadcastPermission).Methods("POST")
	v1.HandleFunc("/publish", api.MemberAuth(api.CreatePublishTopic)).Methods("POST")
	v1.HandleFunc("/publish", api.MemberAuth(api.RemovePublishTopic)).Methods("DELETE")
	v1.HandleFunc("/subscribe", api.MemberAuth(api.CreateSubscribeTopic)).Methods("POST")
	v1.HandleFunc("/subscribe", api.HasBroadcastPermission).Methods("DELETE")
	v1.HandleFunc("/session", api.CreateSession).Methods("POST")

	fmt.Printf("Running on [%s]!\n", viper.GetString("app.address"))
	http.ListenAndServe(viper.GetString("app.address"), v1)
}
