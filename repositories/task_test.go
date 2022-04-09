package repositories

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"fmt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"testing"
)

func CreateTaskMock() (*TaskRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gorm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	if err != nil {
		db.Close()
		return nil, nil, err
	}

	repoTask := MakeTaskRepository(gorm)
	return repoTask, mock, err
}

func TestSelectByIdTask(t *testing.T) {

	var elemID uint = 1

	//создание мока
	repoTask, mock, err := CreateTaskMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_t", "title",
			"description", "position",
			"date_created", "id_l", "id_b"})

	expect := []*models.Task{
		{IdT: elemID, Title: "title", Description: "", Position: 0, DateCreated: "", IdL: 1, IdB: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdT, item.Title, item.Description, item.Position, item.DateCreated, item.IdL, item.IdB)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE "tasks"."id_t" = $1`)).
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repoTask.GetById(elemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE "tasks"."id_t" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrListNotFound)

	_, err = repoTask.GetById(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestCreateTask(t *testing.T) {
	repoTask, mock, err := CreateTaskMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат

	task := models.Task{IdT: 1, Title: "title", Position: 1, IdL: 1, IdB: 1}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "tasks" WHERE id_l = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks" ("title","description","position","date_created","id_l","id_b","id_t") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id_t"`)).
		WithArgs(
			task.Title,
			task.Description,
			task.Position,
			task.DateCreated,
			task.IdL,
			task.IdB,
			task.IdT).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectCommit()

	id, err := repoTask.Create(&task, 1, 1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if id != 1 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// ошибка
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "tasks" WHERE id_l = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks" ("title","description","position","date_created","id_l","id_b","id_t") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id_t"`)).
		WithArgs(
			task.Title,
			task.Description,
			task.Position,
			task.DateCreated,
			task.IdL,
			task.IdB,
			task.IdT).
		WillReturnError(fmt.Errorf("bad_result"))
	mock.ExpectRollback()

	_, err = repoTask.Create(&task, 1, 1)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTasks(t *testing.T) {
	var elemID uint = 1

	//создание мока
	repoTask, mock, err := CreateTaskMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_t", "title", "description", "position",
			"date_created", "id_l", "id_b"})

	expect := []models.Task{
		{IdT: 1, Title: "title", Description: "", Position: 0, DateCreated: "", IdL: 1, IdB: 1},
		{IdT: 2, Title: "title", Description: "", Position: 0, DateCreated: "", IdL: 1, IdB: 1},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdT, item.Title, item.Description, item.Position, item.DateCreated, item.IdL, item.IdB)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE id_l = $1`)).
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repoTask.GetTasks(elemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, &expect) {
		t.Errorf("results not match, want %v, have %v", &expect, item)
		return
	}

	// айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE id_l = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrListNotFound)

	_, err = repoTask.GetTasks(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestDeleteTask(t *testing.T) {
	repoTask, mock, err := CreateTaskMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}
	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_t", "title",
			"description", "position",
			"date_created", "id_l", "id_b"})

	expect := []*models.Task{
		{IdT: 1, Title: "title", Description: "", Position: 0, DateCreated: "", IdL: 1, IdB: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdT, item.Title, item.Description, item.Position, item.DateCreated, item.IdL, item.IdB)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE "tasks"."id_t" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "tasks" WHERE "tasks"."id_t" = $1`)).
		WithArgs(
			1,
		).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "tasks" SET "position"=position - 1 WHERE position > $1 AND id_l = $2`)).
		WithArgs(expect[0].Position,
			expect[0].IdL).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoTask.Delete(1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// ошибка
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE "tasks"."id_t" = $1`)).
		WithArgs(1).
		WillReturnError(customErrors.ErrTaskNotFound)

	err = repoTask.Delete(1)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
