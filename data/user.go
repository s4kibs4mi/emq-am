package data

const (
	UserTypeAdmin   = "admin"
	UserTypeMember  = "member"
	UserTypeDefault = "default"
)

type User struct {
	UserName        string   `json:"user_name"`
	Password        string   `json:"password"`
	Email           string   `json:"email"`
	PublishTopics   []string `json:"publish_topics"`
	SubscribeTopics []string `json:"subscribe_topics"`
	Type            string   `json:"type"`
}

func (u *User) Save() bool {
	return false
}

func (u *User) Find() bool {
	return false
}

func (u *User) Delete() bool {
	return false
}

func (u *User) Count() int {
	return 0
}

func (u *User) ChangePassword() bool {
	return false
}

func (u *User) ChangeUserAccessLevel() bool {
	return false
}

func (u *User) IsUserNameAvailable() bool {
	return false
}

func (u *User) IsEmailAvailable() bool {
	return false
}

func (u *User) HasValidCredentials() bool {
	return false
}

func (u *User) HasPublishPermission(topics []string) bool {
	return false
}

func (u *User) HasSubscribePermission(topics []string) bool {
	return false
}

func (u *User) AppendPublishPermission(topics []string) bool {
	return false
}

func (u *User) AppendSubscribePermission(topics []string) bool {
	return false
}

func (u *User) DiscardPublishPermission(topics []string) bool {
	return false
}

func (u *User) DiscardSubscribePermission(topics []string) bool {
	return false
}
