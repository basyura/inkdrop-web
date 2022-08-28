package models

type Note struct {
	Doctype           string   `json:"doctype"`
	BookId            string   `json:"bookId"`
	CreatedAt         int64    `json:"createdAt"`
	UpdatedAt         int64    `json:"updatedAt"`
	Status            string   `json:"status"`
	Share             string   `json:"share"`
	NumOfTasks        int      `json:"numOfTasks"`
	NumOfCheckedTasks int      `json:"numOfCheckedTasks"`
	Pinned            bool     `json:"pinned"`
	Title             string   `json:"title"`
	Body              string   `json:"body"`
	Tags              []string `json:"tags"`
	Id                string   `json:"id"`
	Rev               string   `json:"rev"`
}
