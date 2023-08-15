package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	DueDate     string
	Priority    int
	Status      string
}

type TaskDAO struct {
	db *gorm.DB
}

func NewTaskDAO(db *gorm.DB) *TaskDAO {
	return &TaskDAO{db: db}
}

func (dao *TaskDAO) Create(task *Task) error {
	return dao.db.Create(task).Error
}

func (dao *TaskDAO) Update(task *Task) error {
	return dao.db.Save(task).Error
}

func (dao *TaskDAO) Delete(task *Task) error {
	return dao.db.Delete(task).Error
}

func (dao *TaskDAO) FindByID(id uint) (*Task, error) {
	var task Task
	err := dao.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (dao *TaskDAO) FindAll() ([]Task, error) {
	var tasks []Task
	err := dao.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
