package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"testing"
)

func CreateCommentMock() (repositories.CommentRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	openGorm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	if err != nil {
		db.Close()
		return nil, nil, err
	}

	repoComment := MakeCommentRepository(openGorm)
	return repoComment, mock, err
}

func TestSelectByIdComment(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoComment, mock, err := CreateCommentMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_cm", "text", "date_created", "id_t", "id_u"})

	expect := []*models.Comment{
		{IdCm: elemID, Text: "text", DateCreated: "", IdT: 1, IdU: 1},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdCm, item.Text, item.DateCreated, item.IdT, item.IdU)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE "comments"."id_cm" = $1`)).
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repoComment.GetById(elemID)
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

	//айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE "comments"."id_cm" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrCommentNotFound)

	_, err = repoComment.GetById(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

//
//func TestCreateCheckListItem(t *testing.T) {
//	t.Parallel()
//
//	repoCheckListItem, mock, err := CreateCheckListItemMock()
//	if err != nil {
//		t.Errorf("unexpected err: %s", err)
//	}
//
//	// нормальный результат
//
//	checkList := models.CheckListItem{IdCl: 1, Description: "title", IdT: 1}
//
//	mock.ExpectBegin()
//	mock.
//		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "check_list_items" ("description","id_cl","id_t","is_ready") VALUES ($1,$2,$3,$4) RETURNING "id_cl_it"`)).
//		WithArgs(
//			checkList.Description,
//			checkList.IdCl,
//			checkList.IdT,
//			checkList.IsReady).
//		WillReturnRows(sqlmock.NewRows([]string{"1"}))
//	mock.ExpectCommit()
//
//	id, err := repoCheckListItem.Create(&checkList)
//	if err != nil {
//		t.Errorf("unexpected err: %s", err)
//		return
//	}
//	if id != 0 {
//		t.Errorf("bad id: want %v, have %v", id, 1)
//		return
//	}
//
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//
//	//ошибка
//	mock.ExpectBegin()
//	mock.
//		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "check_list_items" ("description","id_cl","id_t","is_ready") VALUES ($1,$2,$3,$4) RETURNING "id_cl_it"`)).
//		WithArgs(
//			checkList.Description,
//			checkList.IdCl,
//			checkList.IdT,
//			checkList.IsReady).
//		WillReturnError(fmt.Errorf("bad_result"))
//	mock.ExpectRollback()
//
//	_, err = repoCheckListItem.Create(&checkList)
//	if err == nil {
//		t.Errorf("expected error, got nil")
//		return
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}
//
//func TestDeleteCheckListItem(t *testing.T) {
//	t.Parallel()
//
//	repoCheckListItem, mock, err := CreateCheckListItemMock()
//	if err != nil {
//		t.Errorf("unexpected err: %s", err)
//	}
//	// нормальный результат
//	rows := sqlmock.
//		NewRows([]string{"id_cl_it", "description", "id_cl", "id_t"})
//
//	expect := []*models.CheckListItem{
//		{IdClIt: 1, Description: "description", IdCl: 1, IdT: 1},
//	}
//
//	for _, item := range expect {
//		rows = rows.AddRow(item.IdClIt, item.Description, item.IdT, item.IdCl)
//	}
//
//	mock.ExpectBegin()
//	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "check_list_items" WHERE "check_list_items"."id_cl_it" = $1`)).
//		WithArgs(
//			1,
//		).WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectCommit()
//
//	err = repoCheckListItem.Delete(1)
//	if err != nil {
//		t.Errorf("unexpected err: %s", err)
//		return
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//
//	// ошибка
//	mock.ExpectBegin()
//	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "check_list_items" WHERE "check_list_items"."id_cl_it" = $1`)).
//		WithArgs(
//			1,
//		).WillReturnError(fmt.Errorf("bad_result"))
//	mock.ExpectRollback()
//
//	err = repoCheckListItem.Delete(1)
//	if err == nil {
//		t.Errorf("expected error, got nil")
//		return
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}
//
//func TestUpdateCheckListItem(t *testing.T) {
//	t.Parallel()
//
//	repoCheckListItem, mock, err := CreateCheckListItemMock()
//	if err != nil {
//		t.Errorf("unexpected err: %s", err)
//	}
//
//	// нормальный результат
//	rows := sqlmock.
//		NewRows([]string{"id_cl_it", "description", "id_cl", "id_t"})
//
//	expect := []*models.CheckListItem{
//		{IdClIt: 1, Description: "description", IdCl: 1, IdT: 1},
//	}
//
//	for _, item := range expect {
//		rows = rows.AddRow(item.IdClIt, item.Description, item.IdT, item.IdCl)
//	}
//
//	mock.
//		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "check_list_items" WHERE "check_list_items"."id_cl_it" = $1`)).
//		WithArgs(1).
//		WillReturnRows(rows)
//	mock.ExpectBegin()
//	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "check_list_items" SET "description"=$1,"id_cl"=$2,"id_t"=$3,"is_ready"=$4 WHERE "id_cl_it" = $5`)).
//		WithArgs(
//			expect[0].Description,
//			expect[0].IdT,
//			expect[0].IdCl,
//			expect[0].IsReady,
//			expect[0].IdClIt).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectCommit()
//
//	err = repoCheckListItem.Update(*(expect)[0])
//	if err != nil {
//		t.Errorf("unexpected err: %s", err)
//		return
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//		return
//	}
//
//	//айдишника не существует
//	mock.
//		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "check_list_items" WHERE "check_list_items"."id_cl_it" = $1`)).
//		WithArgs(2).
//		WillReturnError(customErrors.ErrCheckListItemNotFound)
//
//	err = repoCheckListItem.Update(models.CheckListItem{IdClIt: 2})
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//		return
//	}
//	if err == nil {
//		t.Errorf("expected error, got nil")
//		return
//	}
//}
