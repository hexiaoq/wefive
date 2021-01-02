package model

type Comment struct {
	CommentId int64  `json:"comment_id"`
	UserId    int64  `json:"user_id"`
	DeptId    int64  `json:"dept_id"`
	BusId     int64  `json:"bus_id"`
	Content   string `json:"content"`
	Reply     string `json:"reply"`
	Delete1   string `json:"delete_1"`
}
