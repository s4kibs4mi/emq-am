package data

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/s4kibs4mi/emq-am/net"
)

/**
 * := Coded with love by Sakib Sami on 18/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

type Session struct {
	Id           bson.ObjectId `bson:"_id",json:"id",omitempty`
	UserId       bson.ObjectId `json:"user_id"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	CreatedAt    time.Time     `json:"created_at"`
	ExpireAt     time.Time     `json:"expire_at"`
}

func (s *Session) Save() bool {
	err := net.GetSessionCollection().Insert(s)
	s.Find()
	return err == nil
}

func (s *Session) Find() bool {
	results := net.GetSessionCollection().Find(bson.M{
		"accesstoken": s.AccessToken,
		"userid":      s.UserId,
	})
	err := results.One(s)
	return err == nil
}
