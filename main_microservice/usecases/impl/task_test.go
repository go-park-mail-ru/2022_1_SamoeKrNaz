package impl

import (
	customErrors "PLANEXA_backend/errors"
	mock_repositories "PLANEXA_backend/main_microservice/repositories/mocks"
	"PLANEXA_backend/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTasks(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	taskUseCase := MakeTaskUsecase(taskRepo, boardRepo, listRepo, userRepo, checkListRepo, commentRepo)

	tasks := []models.Task{{IdT: 11, Title: "title1"}, {IdT: 12, Title: "title2"}}
	list := models.List{Title: "title1", Position: 1}

	listRepo.EXPECT().GetById(uint(11)).Return(&list, nil)
	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(0)).Return(true, nil)
	taskRepo.EXPECT().GetTasks(uint(11)).Return(&tasks, nil)
	newTasks, err := taskUseCase.GetTasks(uint(11), uint(22))
	assert.Equal(t, tasks, newTasks)
	assert.Equal(t, err, nil)

	listRepo.EXPECT().GetById(uint(11)).Return(&list, nil)
	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(0)).Return(false, nil)
	newTasks, err = taskUseCase.GetTasks(uint(11), uint(22))
	assert.Equal(t, []models.Task(nil), newTasks)
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

//func TestCreateTask(t *testing.T) {
//	t.Parallel()
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//
//	boardRepo := mock_repositories.NewMockBoardRepository(controller)
//	listRepo := mock_repositories.NewMockListRepository(controller)
//	taskRepo := mock_repositories.NewMockTaskRepository(controller)
//	userRepo := mock_repositories.NewMockUserRepository(controller)
//	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
//	commentRepo := mock_repositories.NewMockCommentRepository(controller)
//	taskUseCase := MakeTaskUsecase(taskRepo, boardRepo, listRepo, userRepo, checkListRepo, commentRepo)
//
//	task := models.Task{IdT: 0, Title: "title2"}
//
//	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(0)).Return(true, nil)
//	taskRepo.EXPECT().Create(gomock.Any(), uint(0), uint(0)).Return(uint(0), nil)
//	taskRepo.EXPECT().GetById(uint(0)).Return(&task, nil)
//	_, err := taskUseCase.CreateTask(task, uint(0), uint(0), uint(22))
//	assert.Equal(t, nil, err)
//}

func TestRefactorTask(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	taskUseCase := MakeTaskUsecase(taskRepo, boardRepo, listRepo, userRepo, checkListRepo, commentRepo)

	task := models.Task{IdT: 12, Title: "title2"}

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(0)).Return(true, nil)
	taskRepo.EXPECT().Update(task).Return(nil)
	err := taskUseCase.RefactorTask(task, uint(22))
	assert.Equal(t, err, nil)

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(0)).Return(false, nil)
	err = taskUseCase.RefactorTask(task, uint(22))
	assert.Equal(t, err, customErrors.ErrNoAccess)

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(0)).Return(false, customErrors.ErrTaskNotFound)
	err = taskUseCase.RefactorTask(task, uint(22))
	assert.Equal(t, err, customErrors.ErrTaskNotFound)
}

func TestDeleteTask(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	taskUseCase := MakeTaskUsecase(taskRepo, boardRepo, listRepo, userRepo, checkListRepo, commentRepo)

	task := models.Task{IdT: 0, Title: "title2"}

	taskRepo.EXPECT().GetById(uint(0)).Return(&task, nil)
	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(0)).Return(true, nil)
	taskRepo.EXPECT().Delete(uint(0)).Return(nil)
	err := taskUseCase.DeleteTask(uint(0), uint(22))
	assert.Equal(t, err, nil)

	taskRepo.EXPECT().GetById(uint(0)).Return(&task, nil)
	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(0)).Return(false, nil)
	err = taskUseCase.DeleteTask(uint(0), uint(22))
	assert.Equal(t, err, customErrors.ErrNoAccess)

	taskRepo.EXPECT().GetById(uint(0)).Return(&task, nil)
	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(0)).Return(false, customErrors.ErrBadInputData)
	err = taskUseCase.DeleteTask(uint(0), uint(22))
	assert.Equal(t, err, customErrors.ErrBadInputData)

	taskRepo.EXPECT().GetById(uint(0)).Return(nil, customErrors.ErrTaskNotFound)
	err = taskUseCase.DeleteTask(uint(0), uint(22))
	assert.Equal(t, err, customErrors.ErrTaskNotFound)
}

func TestGetImportantTask(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	taskUseCase := MakeTaskUsecase(taskRepo, boardRepo, listRepo, userRepo, checkListRepo, commentRepo)

	task := models.Task{IdT: 0, Title: "title2"}
	tasks := []models.Task{task, task}

	taskRepo.EXPECT().GetImportantTasks(uint(22)).Return(&tasks, nil)
	newTasks, err := taskUseCase.GetImportantTask(uint(22))
	assert.Equal(t, err, nil)
	assert.Equal(t, tasks, *newTasks)

	taskRepo.EXPECT().GetImportantTasks(uint(22)).Return(nil, customErrors.ErrTaskNotFound)
	_, err = taskUseCase.GetImportantTask(uint(22))
	assert.Equal(t, err, customErrors.ErrTaskNotFound)
}

func TestAppendUser(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	taskUseCase := MakeTaskUsecase(taskRepo, boardRepo, listRepo, userRepo, checkListRepo, commentRepo)

	user := models.User{IdU: 22, Username: "user1"}

	taskRepo.EXPECT().IsAccessToTask(uint(22), uint(0)).Return(true, nil)
	taskRepo.EXPECT().AppendUser(uint(0), uint(22)).Return(nil)
	userRepo.EXPECT().GetUserById(uint(22)).Return(&user, nil)
	newUser, err := taskUseCase.AppendUserToTask(uint(22), uint(22), uint(0))
	assert.Equal(t, err, nil)
	assert.Equal(t, user, newUser)

	taskRepo.EXPECT().IsAccessToTask(uint(22), uint(0)).Return(false, nil)
	_, err = taskUseCase.AppendUserToTask(uint(22), uint(22), uint(0))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestDeleteUser(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	taskUseCase := MakeTaskUsecase(taskRepo, boardRepo, listRepo, userRepo, checkListRepo, commentRepo)

	taskRepo.EXPECT().IsAccessToTask(uint(22), uint(0)).Return(true, nil)
	taskRepo.EXPECT().DeleteUser(uint(0), uint(22)).Return(nil)
	err := taskUseCase.DeleteUserFromTask(uint(22), uint(22), uint(0))
	assert.Equal(t, err, nil)

	taskRepo.EXPECT().IsAccessToTask(uint(22), uint(0)).Return(false, nil)
	err = taskUseCase.DeleteUserFromTask(uint(22), uint(22), uint(0))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}
