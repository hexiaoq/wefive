package model

type SubChat struct {
	SubId   int64  `json:"sub_id"`
	UserId  int64  `json:"user_id"`
	ChatId  int64  `json:"chat_id"`
	Content string `json:"content"`
	Picture string `json:"picture"`
}

type SubChatDto struct {
	SubId   int64  `json:"sub_id"`
	UserId  int64  `json:"user_id"`
	ChatId  int64  `json:"chat_id"`
	Content string `json:"content"`
	Picture string `json:"picture"`
	Avatar  string `json:"avatar"`
	Name    string `json:"name"`
}
