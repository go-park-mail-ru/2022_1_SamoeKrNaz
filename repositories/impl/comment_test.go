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

func TestCreateComment(t *testing.T) {
	t.Parallel()

	repoComment, mock, err := CreateCommentMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат

	comment := models.Comment{IdCm: 1, Text: "text", DateCreated: "", IdT: 1, IdU: 1}

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "comments" ("text","date_created","id_t","id_u","id_cm") VALUES ($1,$2,$3,$4,$5) RETURNING "id_cm"`)).
		WithArgs(
			comment.Text,
			comment.DateCreated,
			comment.IdT,
			comment.IdU,
			comment.IdCm).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectCommit()

	id, err := repoComment.Create(&comment)
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
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "comments" ("text","date_created","id_t","id_u","id_cm") VALUES ($1,$2,$3,$4,$5) RETURNING "id_cm"`)).
		WithArgs(
			comment.Text,
			comment.DateCreated,
			comment.IdT,
			comment.IdU,
			comment.IdCm).
		WillReturnError(fmt.Errorf("bad_result"))
	mock.ExpectRollback()

	_, err = repoComment.Create(&comment)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteComment(t *testing.T) {
	t.Parallel()

	repoComment, mock, err := CreateCommentMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}
	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_cm", "text", "date_created", "id_t", "id_u"})

	expect := []*models.Comment{
		{IdCm: 1, Text: "text", DateCreated: "", IdT: 1, IdU: 1},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdCm, item.Text, item.DateCreated, item.IdT, item.IdU)
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "comments" WHERE "comments"."id_cm" = $1`)).
		WithArgs(
			1,
		).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoComment.Delete(1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// ошибка
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "comments" WHERE "comments"."id_cm" = $1`)).
		WithArgs(
			1,
		).WillReturnError(fmt.Errorf("bad_result"))
	mock.ExpectRollback()

	err = repoComment.Delete(1)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateComment(t *testing.T) {
	t.Parallel()

	repoComment, mock, err := CreateCommentMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_cm", "text", "date_created", "id_t", "id_u"})

	expect := []*models.Comment{
		{IdCm: 1, Text: "text", DateCreated: "", IdT: 1, IdU: 1},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdCm, item.Text, item.DateCreated, item.IdT, item.IdU)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE "comments"."id_cm" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "comments" SET "text"=$1,"date_created"=$2,"id_t"=$3,"id_u"=$4 WHERE "id_cm" = $5`)).
		WithArgs(
			expect[0].Text,
			expect[0].DateCreated,
			expect[0].IdT,
			expect[0].IdU,
			expect[0].IdCm).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoComment.Update(*(expect)[0])
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE "comments"."id_cm" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrCheckListItemNotFound)

	err = repoComment.Update(models.Comment{IdCm: 2})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetComments(t *testing.T) {
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
		{IdCm: 1, Text: "text", DateCreated: "", IdT: 1, IdU: 1},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdCm, item.Text, item.DateCreated, item.IdT, item.IdU)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE id_t = $1 ORDER BY id_cm`)).
		WithArgs(1).
		WillReturnRows(rows)

	item, err := repoComment.GetComments(elemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(&(*item)[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], &(*item)[0])
		return
	}

	// айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE id_t = $1 ORDER BY id_cm`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrBoardNotFound)

	_, err = repoComment.GetComments(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestIsAccessComment(t *testing.T) {
	t.Parallel()

	//создание мока
	repoComment, mock, err := CreateCommentMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_cm", "text", "date_created", "id_t", "id_u"})

	expect := []*models.Comment{
		{IdCm: 1, Text: "text", DateCreated: "", IdT: 1, IdU: 1},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.IdCm, item.Text, item.DateCreated, item.IdT, item.IdU)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE id_cm = $1 and id_u = $2`)).
		WithArgs(1, 1).
		WillReturnRows(rows)

	item, err := repoComment.IsAccessToComment(1, 1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, true) {
		t.Errorf("results not match, want %v, have %v", expect[0], true)
		return
	}

	// айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE id_cm = $1 and id_u = $2`)).
		WithArgs(1, 2).
		WillReturnError(customErrors.ErrCommentNotFound)

	_, err = repoComment.IsAccessToComment(1, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
