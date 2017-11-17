package cmd

import (
	"github.com/spf13/cobra"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/spf13/viper"
)

var ServeCmd = cobra.Command{
	Use:   "serve",
	Run:   ServeCmdExecute,
	Short: "Start server",
}

func ServeCmdExecute(command *cobra.Command, args []string) {
	router := mux.NewRouter()
	v1 := router.PathPrefix("/api/v1").Subrouter()

	fmt.Println("Running on [", viper.GetString("app.address"), "]!")
	http.ListenAndServe(viper.GetString("app.address"), v1)
}
