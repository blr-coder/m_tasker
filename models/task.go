package models

type Task struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Done        bool   `json:"done" bson:"done"`
}

type Tasks []Task

type TaskFilter struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        *bool  `json:"done"`
}

type TaskDelete struct {
	DeletedCount int64 `json:"deleted_count"`
}
