package proto

type UserEnterRoom struct {
	UserId int64 `json:"user_id"`
	UserName string `json:"username"`
}