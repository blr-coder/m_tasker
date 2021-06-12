package interfaces

import "github.com/blr-coder/m_tasker/models"

type TaskInterface interface {
	Create(task models.Task) (models.Task, error)
	Get(id string) (models.Task, error)
	Delete(id string) (models.Task, error)
}