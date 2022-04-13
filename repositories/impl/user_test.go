package impl

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

func CreateUserMock() (*UserRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	openGorm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	if err != nil {
		db.Close()
		return nil, nil, err
	}

	repoUser := MakeUserRepository(openGorm)
	return repoUser, mock, err
}

func TestSelectByIdUser(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoUser, mock, err := CreateUserMock()
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
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repoUser.GetUserById(elemID)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id_u" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrBoardNotFound)

	_, err = repoUser.GetUserById(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	repoUser, mock, err := CreateUserMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	user := models.User{IdU: 1, Username: "user", Password: "", ImgAvatar: ""}

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("username","password","img_avatar","id_u") VALUES ($1,$2,$3,$4) RETURNING "id_u"`)).
		WithArgs(
			user.Username,
			user.Password,
			user.ImgAvatar,
			user.IdU).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectCommit()

	_, err = repoUser.Create(&user)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// ошибка

	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("username","password","img_avatar","id_u") VALUES ($1,$2,$3,$4) RETURNING "id_u"`)).
		WithArgs(
			user.Username,
			user.Password,
			user.ImgAvatar,
			user.IdU).
		WillReturnError(fmt.Errorf("bad_result"))
	mock.ExpectRollback()

	_, err = repoUser.Create(&user)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByLogin(t *testing.T) {
	t.Parallel()

	//создание мока
	repoUser, mock, err := CreateUserMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_u", "username", "password", "img_avatar"})

	expect := []*models.User{
		{IdU: 1, Username: "user", Password: "", ImgAvatar: ""},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdU, item.Username, item.Password, item.ImgAvatar)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE username = $1`)).
		WithArgs("user").
		WillReturnRows(rows)

	item, err := repoUser.GetUserByLogin("user")
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

	//  не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE username = $1`)).
		WithArgs("user2").
		WillReturnError(customErrors.ErrBoardNotFound)

	_, err = repoUser.GetUserByLogin("user2")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestUpdateUser(t *testing.T) {
	t.Parallel()

	repoUser, mock, err := CreateUserMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_u", "username", "password", "img_avatar"})

	expect := []*models.User{
		{IdU: 1, Username: "user", Password: "", ImgAvatar: ""},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdU, item.Username, item.Password, item.ImgAvatar)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id_u" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE username = $1`)).
		WithArgs("newuser").
		WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "username"=$1,"password"=$2,"img_avatar"=$3 WHERE "id_u" = $4`)).
		WithArgs(
			"newuser",
			expect[0].Password,
			expect[0].ImgAvatar,
			expect[0].IdU).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoUser.Update(&models.User{IdU: 1, Username: "newuser", Password: "", ImgAvatar: ""})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}
