package model

type Chat struct {
	ChatId      int64  `json:"chat_id"`
	Title       string `json:"title"`
	UserId      int64  `json:"user_id"`
	Content     string `json:"content"`
	Discussions int    `json:"discussions"`
	Likes       int    `json:"likes"`
	Picture     string `json:"picture"`
}

type ChatDto struct {
	ChatId      int64  `json:"chat_id"`
	Content     string `json:"content"`
	Discussions int    `json:"discussions"`
	Likes       int    `json:"likes"`
	Picture     string `json:"picture"`
	UserId      int64  `json:"user_id"`
	Avatar      string `json:"avatar"`
	Name        string `json:"name"`
}
