package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/models"
	"fmt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"testing"
)

func CreateBoardMock() (repositories.BoardRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	openGorm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	if err != nil {
		db.Close()
		return nil, nil, err
	}

	repoBoard := MakeBoardRepository(openGorm)
	return repoBoard, mock, err
}

func TestSelectByIdBoard(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoBoard, mock, err := CreateBoardMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_b", "title", "description", "img_desk", "data_created", "id_u"})

	expect := []*models.Board{
		{IdB: elemID, Title: "title", Description: "", ImgDesk: "",
			DateCreated: "", IdU: elemID},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdB, item.Title, item.Description,
			item.ImgDesk, item.DateCreated, item.IdU)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boards" WHERE "boards"."id_b" = $1`)).
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repoBoard.GetById(elemID)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boards" WHERE "boards"."id_b" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrBoardNotFound)

	_, err = repoBoard.GetById(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestCreateBoard(t *testing.T) {
	t.Parallel()

	repoBoard, mock, err := CreateBoardMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат

	board := models.Board{IdB: 1, Title: "title", Description: "", ImgDesk: "", DateCreated: ""}

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "boards" ("title","description","img_desk","date_created","id_u","link","id_b") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id_b"`)).
		WithArgs(
			board.Title,
			board.Description,
			board.ImgDesk,
			board.DateCreated,
			board.IdU,
			board.Link,
			board.IdB).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectCommit()

	id, err := repoBoard.Create(&board)
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

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "boards" ("title","description","img_desk","date_created","id_u","link","id_b") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id_b"`)).
		WithArgs(
			board.Title,
			board.Description,
			board.ImgDesk,
			board.DateCreated,
			board.IdU,
			board.Link,
			board.IdB).
		WillReturnError(fmt.Errorf("bad_result"))
	mock.ExpectRollback()

	_, err = repoBoard.Create(&board)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteBoard(t *testing.T) {
	t.Parallel()
	repoBoard, mock, err := CreateBoardMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}
	// нормальный результат
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users_boards" WHERE "users_boards"."board_id_b" = $1`)).
		WithArgs(
			1,
		).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "boards" WHERE "boards"."id_b" = $1`)).
		WithArgs(
			1,
		).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoBoard.Delete(1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// ошибка
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users_boards" WHERE "users_boards"."board_id_b" = $1`)).
		WithArgs(
			1,
		).WillReturnError(fmt.Errorf("bad_result"))
	mock.ExpectRollback()

	err = repoBoard.Delete(1)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateBoard(t *testing.T) {
	t.Parallel()

	repoBoard, mock, err := CreateBoardMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_b", "title", "description", "img_desk", "data_created", "id_u"})

	expect := []*models.Board{
		{IdB: 1, Title: "title", Description: "", ImgDesk: "",
			DateCreated: "", IdU: 1},
		{IdB: 1, Title: "title2", Description: "descr", ImgDesk: "",
			DateCreated: "", IdU: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdB, item.Title, item.Description,
			item.ImgDesk, item.DateCreated, item.IdU)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boards" WHERE "boards"."id_b" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "boards" SET "title"=$1,"description"=$2,"img_desk"=$3,"date_created"=$4,"id_u"=$5,"link"=$6 WHERE "id_b" = $7`)).
		WithArgs(expect[1].Title,
			expect[1].Description,
			expect[1].ImgDesk,
			expect[1].DateCreated,
			expect[1].IdU,
			expect[1].Link,
			expect[1].IdB).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoBoard.Update(*(expect)[1])
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boards" WHERE "boards"."id_b" = $1`)).
		WithArgs(3).
		WillReturnError(customErrors.ErrBoardNotFound)

	err = repoBoard.Update(models.Board{IdB: 3})
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestAppendUser(t *testing.T) {
	t.Parallel()

	//создание мока
	repoBoard, mock, err := CreateBoardMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_b", "title", "description", "img_desk", "data_created", "id_u"})

	expect := []*models.Board{
		{IdB: 1, Title: "title", Description: "", ImgDesk: "",
			DateCreated: "", IdU: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdB, item.Title, item.Description,
			item.ImgDesk, item.DateCreated, item.IdU)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id_u" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)
	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("username","password","img_avatar","id_u") VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING RETURNING "id_u"`)).
		WithArgs("", "", "", 1).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.
		ExpectExec(regexp.QuoteMeta(`INSERT INTO "users_boards" ("board_id_b","user_id_u") VALUES ($1,$2) ON CONFLICT DO NOTHING`)).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoBoard.AppendUser(1, 1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id_u" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrBoardNotFound)

	err = repoBoard.AppendUser(2, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetLists(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoBoard, mock, err := CreateBoardMock()
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "lists" WHERE id_b = $1 ORDER BY position`)).
		WithArgs(1).
		WillReturnRows(rows)

	item, err := repoBoard.GetLists(elemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(&item[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], &item[0])
		return
	}

	// айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "lists" WHERE id_b = $1 ORDER BY position`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrBoardNotFound)

	_, err = repoBoard.GetLists(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetUserBoard(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoBoard, mock, err := CreateBoardMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_b", "title", "description", "img_desk", "data_created", "id_u"})

	expect := []*models.Board{
		{IdB: elemID, Title: "title", Description: "", ImgDesk: "",
			DateCreated: "", IdU: elemID},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdB, item.Title, item.Description,
			item.ImgDesk, item.DateCreated, item.IdU)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT "boards"."id_b","boards"."title","boards"."description","boards"."img_desk","boards"."date_created","boards"."id_u","boards"."link" FROM "boards" JOIN "users_boards" ON "users_boards"."board_id_b" = "boards"."id_b" AND "users_boards"."user_id_u" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	item, err := repoBoard.GetUserBoards(elemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(&item[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], &item[0])
		return
	}

	// айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT "boards"."id_b","boards"."title","boards"."description","boards"."img_desk","boards"."date_created","boards"."id_u","boards"."link" FROM "boards" JOIN "users_boards" ON "users_boards"."board_id_b" = "boards"."id_b" AND "users_boards"."user_id_u" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrBoardNotFound)

	_, err = repoBoard.GetUserBoards(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestIsAccess(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoBoard, mock, err := CreateBoardMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_b", "title", "description", "img_desk", "data_created", "id_u"})

	expect := []*models.Board{
		{IdB: elemID, Title: "title", Description: "", ImgDesk: "",
			DateCreated: "", IdU: elemID},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdB, item.Title, item.Description,
			item.ImgDesk, item.DateCreated, item.IdU)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT "users"."id_u","users"."username","users"."password","users"."img_avatar" FROM "users" JOIN "users_boards" ON "users_boards"."user_id_u" = "users"."id_u" AND "users_boards"."board_id_b" = $1 WHERE id_u = $2`)).
		WithArgs(1, 1).
		WillReturnRows(rows)

	item, err := repoBoard.IsAccessToBoard(1, 1)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT "users"."id_u","users"."username","users"."password","users"."img_avatar" FROM "users" JOIN "users_boards" ON "users_boards"."user_id_u" = "users"."id_u" AND "users_boards"."board_id_b" = $1 WHERE id_u = $2`)).
		WithArgs(1, 1).
		WillReturnError(customErrors.ErrBoardNotFound)

	_, err = repoBoard.IsAccessToBoard(1, 1)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetBoardUser(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoBoard, mock, err := CreateBoardMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_u", "username", "password", "img_avatar"})

	expect := []*models.User{
		{IdU: elemID, Username: "user", Password: "", ImgAvatar: ""},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdU, item.Username, item.Password, item.ImgAvatar)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT "users"."id_u","users"."username","users"."password","users"."img_avatar" FROM "users" JOIN "users_boards" ON "users_boards"."user_id_u" = "users"."id_u" AND "users_boards"."board_id_b" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	item, err := repoBoard.GetBoardUser(1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(&item[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item[0])
		return
	}

	// айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT "users"."id_u","users"."username","users"."password","users"."img_avatar" FROM "users" JOIN "users_boards" ON "users_boards"."user_id_u" = "users"."id_u" AND "users_boards"."board_id_b" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrBoardNotFound)

	_, err = repoBoard.GetBoardUser(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestDeleteUser(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoBoard, mock, err := CreateBoardMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_u", "username", "password", "img_avatar"})

	expect := []*models.User{
		{IdU: elemID, Username: "user", Password: "", ImgAvatar: ""},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdU, item.Username, item.Password, item.ImgAvatar)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id_u" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)
	mock.ExpectBegin()
	mock.
		ExpectExec(regexp.QuoteMeta(`DELETE FROM "users_boards" WHERE "users_boards"."board_id_b" = $1 AND "users_boards"."user_id_u" = $2`)).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoBoard.DeleteUser(1, 1)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id_u" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrUserNotFound)

	err = repoBoard.DeleteUser(1, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
