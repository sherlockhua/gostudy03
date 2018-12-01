package proto


const (
	MessageTypeUserEnterRoom = 1001
	MessageTypeUserTalk      = 1002
	MessageTypeOnlineUserList = 1003
)


type UserEnterRoom struct {
	UserId int64 `json:"user_id"`
	UserName string `json:"username"`
	ImageUrl string `json:"image_url"`
}

type OnlineUser struct{
	UserId int64 `json:"user_id"`
	UserName string `json:"username"`
	ImageUrl string `json:"image_url"`
}

type OnlineUserList struct {
	OnlineUserCount int `json:"online_user_count"`
	OnlineUserArray []*OnlineUser `json:"user_list"`
}

type Message struct {
	Type int `json:"type"`
	Data []byte `json:"data"`
}

type UserTalk struct {
	UserName string `json:"username"`
	UserId int64 `json:"user_id"`
	Content string `json:"content"`
	ImageUrl string `json:"image_url"`
}