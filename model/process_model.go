package model

type Process struct {
	ProcessId int64  `json:"process_id"`
	BusId     int64  `json:"bus_id"`
	Step      int    `json:"step"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type ProcessDto struct {
	ProcessId int64      `json:"process_id"`
	BusId     int64      `json:"bus_id"`
	Step      int        `json:"step"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Materials []Material `json:"materials"`
}
