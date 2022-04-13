package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"fmt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"testing"
)

func CreateListMock() (repositories.ListRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	openGorm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	if err != nil {
		db.Close()
		return nil, nil, err
	}

	repoList := MakeListRepository(openGorm)
	return repoList, mock, err
}

func TestSelectByIdList(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoList, mock, err := CreateListMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_l", "title", "position", "id_b"})

	expect := []*models.List{
		{IdL: elemID, Title: "title", Position: 0, IdB: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdL, item.Title, item.Position, item.IdB)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "lists" WHERE "lists"."id_l" = $1`)).
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repoList.GetById(elemID)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "lists" WHERE "lists"."id_l" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrListNotFound)

	_, err = repoList.GetById(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestCreateList(t *testing.T) {
	t.Parallel()

	repoList, mock, err := CreateListMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат

	list := models.List{IdL: 1, Title: "title", Position: 1, IdB: 1}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "lists" WHERE id_b = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "lists" ("title","position","id_b","id_l") VALUES ($1,$2,$3,$4) RETURNING "id_l"`)).
		WithArgs(
			list.Title,
			list.Position,
			list.IdB,
			list.IdL).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectCommit()

	id, err := repoList.Create(&list, 1)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "lists" WHERE id_b = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "lists" ("title","position","id_b","id_l") VALUES ($1,$2,$3,$4) RETURNING "id_l"`)).
		WithArgs(
			list.Title,
			list.Position,
			list.IdB,
			list.IdL).
		WillReturnError(fmt.Errorf("bad_result"))
	mock.ExpectRollback()

	_, err = repoList.Create(&list, 1)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteList(t *testing.T) {
	t.Parallel()

	repoList, mock, err := CreateListMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}
	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_l", "title", "position", "id_b"})

	expect := []*models.List{
		{IdL: 1, Title: "title", Position: 0, IdB: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdL, item.Title, item.Position, item.IdB)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "lists" WHERE "lists"."id_l" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "lists" WHERE "lists"."id_l" = $1`)).
		WithArgs(
			1,
		).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "lists" SET "position"=position - 1 WHERE position > $1 AND id_b = $2`)).
		WithArgs(expect[0].Position,
			expect[0].IdL).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoList.Delete(1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// ошибка
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "lists" WHERE "lists"."id_l" = $1`)).
		WithArgs(1).
		WillReturnError(customErrors.ErrTaskNotFound)

	err = repoList.Delete(1)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTasksList(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoList, mock, err := CreateListMock()
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

	item, err := repoList.GetTasks(elemID)
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

	_, err = repoList.GetTasks(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetBoard(t *testing.T) {
	t.Parallel()

	//создание мока
	repoList, mock, err := CreateListMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_t", "title", "position", "id_b"})

	expect := []models.List{
		{IdL: 0, Title: "title", Position: 2, IdB: 2},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdL, item.Title, item.Position, item.IdB)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "lists" WHERE "lists"."id_l" = $1`)).
		WithArgs(0).
		WillReturnRows(rows)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boards" WHERE "boards"."id_b" = $1`)).
		WithArgs(2).
		WillReturnRows(rows)

	_, err = repoList.GetBoard(0)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
