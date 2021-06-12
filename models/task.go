package models

type Task struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Done        bool   `json:"done" bson:"done"`
}
