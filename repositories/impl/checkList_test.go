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

func CreateCheckListMock() (repositories.CheckListRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	openGorm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	if err != nil {
		db.Close()
		return nil, nil, err
	}

	repoList := MakeCheckListRepository(openGorm)
	return repoList, mock, err
}

func TestSelectByIdCheckList(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoCheckList, mock, err := CreateCheckListMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_cl", "title", "id_t"})

	expect := []*models.CheckList{
		{IdCl: elemID, Title: "title", IdT: 1},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdCl, item.Title, item.IdT)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "check_lists" WHERE "check_lists"."id_cl" = $1`)).
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repoCheckList.GetById(elemID)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "check_lists" WHERE "check_lists"."id_cl" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrCheckListNotFound)

	_, err = repoCheckList.GetById(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestCreateCheckList(t *testing.T) {
	t.Parallel()

	repoCheckList, mock, err := CreateCheckListMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат

	checkList := models.CheckList{IdCl: 1, Title: "title", IdT: 1}

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "check_lists" ("title","id_t","id_cl") VALUES ($1,$2,$3) RETURNING "id_cl"`)).
		WithArgs(
			checkList.Title,
			checkList.IdT,
			checkList.IdCl).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectCommit()

	id, err := repoCheckList.Create(&checkList)
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

	//ошибка
	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "check_lists" ("title","id_t","id_cl") VALUES ($1,$2,$3) RETURNING "id_cl"`)).
		WithArgs(
			checkList.Title,
			checkList.IdT,
			checkList.IdCl).
		WillReturnError(fmt.Errorf("bad_result"))
	mock.ExpectRollback()

	_, err = repoCheckList.Create(&checkList)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteCheckList(t *testing.T) {
	t.Parallel()

	repoCheckList, mock, err := CreateCheckListMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}
	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_cl", "title", "id_t"})

	expect := []*models.CheckList{
		{IdCl: 1, Title: "title", IdT: 1},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdCl, item.Title, item.IdT)
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "check_lists" WHERE "check_lists"."id_cl" = $1`)).
		WithArgs(
			1,
		).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoCheckList.Delete(1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// ошибка
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "check_lists" WHERE "check_lists"."id_cl" = $1`)).
		WithArgs(
			1,
		).WillReturnError(fmt.Errorf("bad_result"))
	mock.ExpectRollback()

	err = repoCheckList.Delete(1)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetCheckListItems(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoCheckList, mock, err := CreateCheckListMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_cl_it", "description",
			"id_cl", "id_t"})

	expect := []models.CheckListItem{
		{IdClIt: 1, Description: "", IdCl: 1, IdT: 1},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdClIt, item.Description, item.IdCl, item.IdT)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "check_list_items" WHERE id_cl = $1`)).
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repoCheckList.GetCheckListItems(elemID)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "check_list_items" WHERE id_cl = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrListNotFound)

	_, err = repoCheckList.GetCheckListItems(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestUpdateCheckList(t *testing.T) {
	t.Parallel()

	repoCheckList, mock, err := CreateCheckListMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_cl", "title", "id_t"})

	expect := []*models.CheckList{
		{IdCl: 1, Title: "title", IdT: 1},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdCl, item.Title, item.IdT)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "check_lists" WHERE "check_lists"."id_cl" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "check_lists" SET "title"=$1,"id_t"=$2 WHERE "id_cl" = $3`)).
		WithArgs(
			expect[0].Title,
			expect[0].IdT,
			expect[0].IdCl).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoCheckList.Update(*(expect)[0])
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	//айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "check_lists" WHERE "check_lists"."id_cl" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrBoardNotFound)

	err = repoCheckList.Update(models.CheckList{IdCl: 2})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
