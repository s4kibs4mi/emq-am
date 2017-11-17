package net

import (
	"github.com/spf13/viper"
	"fmt"
	"os"
	"gopkg.in/mgo.v2"
)

/**
 * := Coded with love by Sakib Sami on 17/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

var mSession *mgo.Session
var mDatabase *mgo.Database
var mConnectError error

func NewMongoDBConnection() {
	mSession, mConnectError = mgo.Dial(viper.GetString("databases.mongodb.uri"))
	if mConnectError != nil {
		fmt.Printf("Couldn't connect to database [ %s/%s ]", viper.GetString("databases.mongodb.uri"),
			viper.GetString("databases.mongodb.name"))
		os.Exit(-1)
	}
	mDatabase = mSession.DB(viper.GetString("databases.mongodb.name"))
}

func GetMongoDB() *mgo.Database {
	return mDatabase
}

func GetUserCollection() *mgo.Collection {
	return GetMongoDB().C(viper.GetString("databases.mongodb.auth_collection"))
}
