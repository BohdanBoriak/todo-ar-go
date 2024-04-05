package database

import (
	"time"
	"todo-list/domain"

	"github.com/upper/db/v4"
)

const tasksTableName = "tasks"

type task struct {
	Id          uint64            `db:"id,omitempty"`
	UserId      uint64            `db:"user_id"`
	Title       string            `db:"title"`
	Description *string           `db:"description"`
	Status      domain.TaskStatus `db:"status"`
	Date        *time.Time        `db:"date"`
}

type TaskRepository struct {
	coll db.Collection
	sess db.Session
}

func NewTaskRepository(sess db.Session) TaskRepository {
	return TaskRepository{
		coll: sess.Collection(tasksTableName),
		sess: sess,
	}
}

func (r TaskRepository) Save(t domain.Task) (domain.Task, error) {
	tsk := r.mapDomainToModel(t)
	err := r.coll.InsertReturning(&tsk)
	if err != nil {
		return domain.Task{}, err
	}
	newTask := r.mapModelToDomain(tsk)
	return newTask, nil
}

func (r TaskRepository) FindById(id uint64) (domain.Task, error) {
	var t task
	err := r.coll.Find("id = ?", id).One(&t)
	if err != nil {
		return domain.Task{}, err
	}
	tsk := r.mapModelToDomain(t)
	return tsk, nil
}

func (r TaskRepository) FindTasksForToday(uId uint64) ([]domain.Task, error) {
	var tasks []task
	err := r.coll.Find(db.Cond{
		"user_id": uId,
	}).
		And("date = CURRENT_DATE").
		All(&tasks)
	if err != nil {
		return nil, err
	}
	ts := r.mapModelToDomainCollection(tasks)
	return ts, nil
}

func (r TaskRepository) Update(t domain.Task) (domain.Task, error) {
	tsk := r.mapDomainToModel(t)
	err := r.coll.Find("id = ?", t.Id).Update(&tsk)
	if err != nil {
		return domain.Task{}, err
	}
	updatedTask := r.mapModelToDomain(tsk)
	return updatedTask, nil
}

func (r TaskRepository) Delete(id uint64) error {
	err := r.coll.Find("id = ?", id).Delete()
	return err
}

func (r TaskRepository) mapDomainToModel(t domain.Task) task {
	return task{
		Id:          t.Id,
		UserId:      t.UserId,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Date:        t.Date,
	}
}

func (r TaskRepository) mapModelToDomain(t task) domain.Task {
	return domain.Task{
		Id:          t.Id,
		UserId:      t.UserId,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Date:        t.Date,
	}
}

func (r TaskRepository) mapModelToDomainCollection(tasks []task) []domain.Task {
	ts := make([]domain.Task, len(tasks))
	for i, t := range tasks {
		ts[i] = r.mapModelToDomain(t)
	}
	return ts
}
