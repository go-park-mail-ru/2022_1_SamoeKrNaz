package impl

import (
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/models"
	"fmt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func CreateAttachmentMock() (repositories.AttachmentRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	openGorm, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	if err != nil {
		db.Close()
		return nil, nil, err
	}

	repoBoard := MakeAttachmentRepository(openGorm)
	return repoBoard, mock, err
}

func TestGetAttachment(t *testing.T) {
	t.Parallel()

	repoAttach, mock, err := CreateAttachmentMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	// нормальный результат
	rows := sqlmock.
		NewRows([]string{"id_a", "default_name", "system_name", "id_t"})

	expect := []*models.Attachment{
		{IdA: 1, IdT: 1},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.IdA, item.DefaultName, item.SystemName, item.IdT)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "attachments" WHERE "attachments"."id_a" = $1`)).
		WithArgs(
			1).
		WillReturnRows(rows)

	id, err := repoAttach.GetById(1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if *(expect)[0] != *id {
		t.Errorf("bad id: want %v, have %v", id, expect[0])
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	//ошибка

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "attachments" WHERE "attachments"."id_a" = $1`)).
		WithArgs(
			2).
		WillReturnError(fmt.Errorf("bad_result"))

	_, err = repoAttach.GetById(2)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteAttachment(t *testing.T) {
	t.Parallel()
	repoAttach, mock, err := CreateAttachmentMock()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}
	// нормальный результат
	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "attachments" WHERE "attachments"."id_a" = $1`)).
		WithArgs(
			2).
		WillReturnError(fmt.Errorf("bad_result"))

	err = repoAttach.Delete(2)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
