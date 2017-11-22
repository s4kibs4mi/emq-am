package data

import (
	"github.com/s4kibs4mi/emq-am/net"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/s4kibs4mi/emq-am/utils"
	"github.com/spf13/viper"
)

type UserType string
type UserStatus string
type MQTopicDirection string

const (
	Admin   UserType = "admin"
	Member  UserType = "member"
	Default UserType = "default"
)

const (
	Allowed UserStatus = "allowed"
	Blocked UserStatus = "blocked"
)

const (
	Subscribe MQTopicDirection = "1"
	Publish   MQTopicDirection = "2"
)

type ACLParams struct {
	Access MQTopicDirection
	UserId string `json:"user_id"`
	Topic  string `json:"topic"`
}

type UserRequest struct {
	UserName string   `json:"user_name,omitempty"`
	Password string   `json:"password,omitempty"`
	Email    string   `json:"email,omitempty"`
	Type     UserType `json:"type,omitempty"`
}

type User struct {
	Id              bson.ObjectId `json:"id"`
	UserName        string        `json:"user_name,omitempty"`
	Password        string        `json:"-"`
	Email           string        `json:"email,omitempty"`
	PublishTopics   []string      `json:"publish_topics,omitempty"`
	SubscribeTopics []string      `json:"subscribe_topics,omitempty"`
	Type            UserType      `json:"type,omitempty"`
	Status          UserStatus    `json:"status"`
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
	return u.FindByUsername()
}

func (u *User) FindByUsername() bool {
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
	result := net.GetUserCollection().Find(bson.M{
		"id": u.Id,
	})
	err := result.One(u)
	return err == nil
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
		"id":   u.Id,
		"type": Admin,
	})
	n, err := result.Count()
	if err == nil && n == 1 {
		return true
	}
	return false
}

func (u *User) IsMember() bool {
	result := net.GetUserCollection().Find(bson.M{
		"id":   u.Id,
		"type": Member,
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

func (u *User) HasValidCredentials(request *UserRequest) bool {
	result := net.GetUserCollection().Find(bson.M{
		"username": request.UserName,
	})
	err := result.One(u)
	if err == nil {
		return utils.IsPasswordMatched(request.Password, u.Password)
	}
	return false
}

func (u *User) HasPublishPermission(topic string) bool {
	if u.IsAdmin() {
		return true
	}
	result := net.GetUserCollection().Find(bson.M{
		"id":            u.Id,
		"publishtopics": topic,
	})
	n, err := result.Count()
	return err == nil && n == 1
}

func (u *User) HasSubscribePermission(topic string) bool {
	if u.IsAdmin() {
		return true
	}
	result := net.GetUserCollection().Find(bson.M{
		"id":              u.Id,
		"subscribetopics": topic,
	})
	n, err := result.Count()
	return err == nil && n == 1
}

func (u *User) AppendPublishPermission(topic string) bool {
	u.PublishTopics = append(u.PublishTopics, topic)
	err := net.GetUserCollection().Update(bson.M{
		"id": u.Id,
	}, u)
	return err == nil
}

func (u *User) AppendSubscribePermission(topic string) bool {
	u.SubscribeTopics = append(u.SubscribeTopics, topic)
	err := net.GetUserCollection().Update(bson.M{
		"id": u.Id,
	}, u)
	return err == nil
}

func (u *User) DiscardPublishPermission(topic string) bool {
	var newTopicList []string
	for _, t := range u.PublishTopics {
		if t == topic {
			continue
		}
		newTopicList = append(newTopicList, t)
	}
	u.PublishTopics = newTopicList
	err := net.GetUserCollection().Update(bson.M{
		"id": u.Id,
	}, u)
	return err == nil
}

func (u *User) DiscardSubscribePermission(topic string) bool {
	var newTopicList []string
	for _, t := range u.SubscribeTopics {
		if t == topic {
			continue
		}
		newTopicList = append(newTopicList, t)
	}
	u.SubscribeTopics = newTopicList
	err := net.GetUserCollection().Update(bson.M{
		"id": u.Id,
	}, u)
	return err == nil
}

func (u *User) GetUserList(page int) []User {
	users := []User{}
	perPage := viper.GetInt("pagination.per_page")
	result := net.GetUserCollection().Find(bson.M{}).Limit(perPage).Skip(perPage * page)
	result.All(&users)
	return users
}
