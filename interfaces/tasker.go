package interfaces

import "github.com/blr-coder/m_tasker/models"

type TaskInterface interface {
	Create(task models.Task) (models.Task, error)
	Get(id string) (models.Task, error)
	List() (models.Tasks, error)
	Search(filter *models.TaskFilter) (models.Tasks, error)
	Done(id string) (models.Task, error)
	Update(id string, task *models.Task) (models.Task, error)
	Delete(id string) (models.TaskDelete, error)
}
