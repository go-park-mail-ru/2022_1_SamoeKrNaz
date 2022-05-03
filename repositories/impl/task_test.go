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

func CreateTaskMock() (repositories.TaskRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	openGorm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	if err != nil {
		db.Close()
		return nil, nil, err
	}

	repoTask := MakeTaskRepository(openGorm)
	return repoTask, mock, err
}

func TestSelectByIdTask(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	repoTask, mock, err := CreateTaskMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат

	task := models.Task{IdT: 1, Title: "title", Position: 1, IdL: 1, IdB: 1, IdU: 1}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "tasks" WHERE id_l = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks" ("title","description","position","date_created","id_l","id_b","date_to_order","deadline","id_u","is_ready","is_important","id_t") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING "id_t"`)).
		WithArgs(
			task.Title,
			task.Description,
			task.Position,
			task.DateCreated,
			task.IdL,
			task.IdB,
			task.DateToOrder,
			task.Deadline,
			task.IdU,
			task.IsReady,
			task.IsImportant,
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
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks" ("title","description","position","date_created","id_l","id_b","date_to_order","deadline","id_u","is_ready","is_important","id_t") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING "id_t"`)).
		WithArgs(
			task.Title,
			task.Description,
			task.Position,
			task.DateCreated,
			task.IdL,
			task.IdB,
			task.DateToOrder,
			task.Deadline,
			task.IdU,
			task.IsReady,
			task.IsImportant,
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
	t.Parallel()

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
	t.Parallel()

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

func TestAppendUserTask(t *testing.T) {
	t.Parallel()

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
		{IdT: 1, Title: "title", Description: "", Position: 0, DateCreated: "", IdL: 1, IdB: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdT, item.Title, item.Description, item.Position, item.DateCreated, item.IdL, item.IdB)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id_u" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)
	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("username","password","img_avatar") VALUES ($1,$2,$3) ON CONFLICT DO NOTHING RETURNING "id_u"`)).
		WithArgs("", "", "").
		WillReturnRows(sqlmock.NewRows([]string{"1"}))
	mock.
		ExpectExec(regexp.QuoteMeta(`INSERT INTO "users_tasks" ("task_id_t","user_id_u") VALUES ($1,$2) ON CONFLICT DO NOTHING`)).
		WithArgs(1, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoTask.AppendUser(1, 1)
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

	err = repoTask.AppendUser(2, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetCheckLists(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoBoard, mock, err := CreateTaskMock()
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "check_lists" WHERE id_t = $1 ORDER BY id_cl`)).
		WithArgs(1).
		WillReturnRows(rows)

	item, err := repoBoard.GetCheckLists(elemID)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "check_lists" WHERE id_t = $1 ORDER BY id_cl`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrBoardNotFound)

	_, err = repoBoard.GetCheckLists(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestIsAccessTask(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoTask, mock, err := CreateTaskMock()
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE "tasks"."id_t" = $1`)).
		WithArgs(elemID).
		WillReturnRows(rows)
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT "boards"."id_b","boards"."title","boards"."description","boards"."img_desk","boards"."date_created","boards"."id_u" FROM "boards" JOIN "users_boards" ON "users_boards"."board_id_b" = "boards"."id_b" AND "users_boards"."user_id_u" = $1 WHERE id_b = $2`)).
		WithArgs(1, 0).
		WillReturnRows(rows)

	item, err := repoTask.IsAccessToTask(1, 1)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE "tasks"."id_t" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrTaskNotFound)

	_, err = repoTask.IsAccessToTask(1, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetImportantTasks(t *testing.T) {
	t.Parallel()

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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE id_u = $1 and is_important = true ORDER BY date_to_order`)).
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repoTask.GetImportantTasks(elemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(&(*item)[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// айдишника не существует
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE id_u = $1 and is_important = true ORDER BY date_to_order`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrListNotFound)

	_, err = repoTask.GetImportantTasks(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetTaskUser(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoTask, mock, err := CreateTaskMock()
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
		ExpectQuery(regexp.QuoteMeta(`SELECT "users"."id_u","users"."username","users"."password","users"."img_avatar" FROM "users" JOIN "users_tasks" ON "users_tasks"."user_id_u" = "users"."id_u" AND "users_tasks"."task_id_t" = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	item, err := repoTask.GetTaskUser(1)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT "users"."id_u","users"."username","users"."password","users"."img_avatar" FROM "users" JOIN "users_tasks" ON "users_tasks"."user_id_u" = "users"."id_u" AND "users_tasks"."task_id_t" = $1`)).
		WithArgs(2).
		WillReturnError(customErrors.ErrBoardNotFound)

	_, err = repoTask.GetTaskUser(2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestDeleteUserTask(t *testing.T) {
	t.Parallel()

	var elemID uint = 1

	//создание мока
	repoTask, mock, err := CreateTaskMock()
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
		ExpectExec(regexp.QuoteMeta(`DELETE FROM "users_tasks" WHERE "users_tasks"."task_id_t" = $1 AND "users_tasks"."user_id_u" = $2`)).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repoTask.DeleteUser(1, 1)
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

	err = repoTask.DeleteUser(1, 2)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
