package entity

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Vip      int64  `json:"vip"`
}

func (User) TableName() string {
	return "member_info"
}
