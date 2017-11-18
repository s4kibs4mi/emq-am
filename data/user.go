package data

import (
	"github.com/s4kibs4mi/emq-am/net"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/s4kibs4mi/emq-am/utils"
	"fmt"
)

const (
	UserTypeAdmin   = "admin"
	UserTypeMember  = "member"
	UserTypeDefault = "default"
)

const (
	MQTopicDirectionSubscribe = "1"
	MQTopicDirectionPublish   = "2"
)

//access= %A, username = %u, clientid= %c, ipaddr = %a, topic = %t
type ACLParams struct {
	Access   string
	UserName string
	Topic    string `json:"topic"`
}

type User struct {
	Id              bson.ObjectId `bson:"_id,omitempty",json:"id"`
	UserName        string        `json:"user_name,omitempty"`
	Password        string        `json:"password,omitempty"`
	Email           string        `json:"email,omitempty"`
	PublishTopics   []string      `json:"publish_topics,omitempty"`
	SubscribeTopics []string      `json:"subscribe_topics,omitempty"`
	Type            string        `json:"type,omitempty"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

func (u *User) Save() bool {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	err := net.GetUserCollection().Insert(u)
	if err != nil {
		return false
	}
	return u.Find()
}

func (u *User) Find() bool {
	result := net.GetUserCollection().Find(bson.M{
		"username": u.UserName,
	})
	err := result.One(u)
	if err == nil {
		return true
	}
	return false
}

func (u *User) FindById() bool {
	result := net.GetUserCollection().FindId(bson.M{
		"_id": u.Id,
	})
	err := result.One(u)
	if err == nil {
		return true
	}
	return false
}

func (u *User) Delete() bool {
	return false
}

func (u *User) Count() int {
	result := net.GetUserCollection().Find(bson.M{})
	n, err := result.Count()
	if err == nil {
		return n
	}
	return -1
}

func (u *User) ChangePassword() bool {
	return false
}

func (u *User) ChangeUserAccessLevel() bool {
	return false
}

func (u *User) IsUserNameAvailable() bool {
	result := net.GetUserCollection().Find(bson.M{
		"username": u.UserName,
	})
	n, err := result.Count()
	if err == nil && n == 0 {
		return true
	}
	return false
}

func (u *User) IsAdmin() bool {
	result := net.GetUserCollection().Find(bson.M{
		"username": u.UserName,
		"type":     UserTypeAdmin,
	})
	n, err := result.Count()
	if err == nil && n == 1 {
		return true
	}
	return false
}

func (u *User) IsMember() bool {
	result := net.GetUserCollection().Find(bson.M{
		"username": u.UserName,
		"type":     UserTypeMember,
	})
	n, err := result.Count()
	if err == nil && n == 1 {
		return true
	}
	return false
}

func (u *User) IsEmailAvailable() bool {
	result := net.GetUserCollection().Find(bson.M{
		"email": u.Email,
	})
	n, err := result.Count()
	if err == nil && n == 0 {
		return true
	}
	return false
}

func (u *User) HasValidCredentials() bool {
	user := &User{}
	result := net.GetUserCollection().Find(bson.M{
		"username": u.UserName,
	})
	err := result.One(user)
	if err == nil {
		u.Id = user.Id
		return utils.IsPasswordMatched(u.Password, user.Password)
	}
	return false
}

func (u *User) HasPublishPermission(topic string) bool {
	if u.IsAdmin() {
		return true
	}
	result := net.GetUserCollection().Find(bson.M{
		"username":      u.UserName,
		"publishtopics": topic,
	})
	n, err := result.Count()
	if err == nil && n == 1 {
		fmt.Println("true")
		return true
	}
	fmt.Println("false")
	return false
}

func (u *User) HasSubscribePermission(topics string) bool {
	return true
}

func (u *User) AppendPublishPermission(topics string) bool {

	return false
}

func (u *User) AppendSubscribePermission(topics string) bool {
	return false
}

func (u *User) DiscardPublishPermission(topics string) bool {
	return false
}

func (u *User) DiscardSubscribePermission(topics string) bool {
	return false
}
