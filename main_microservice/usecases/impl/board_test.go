package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories/mocks"
	"PLANEXA_backend/models"
	"github.com/golang/mock/gomock"
	rtime "github.com/ivahaev/russian-time"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"strconv"
	"testing"
	"time"
)

func TestGetBoards(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo, commentRepo)

	boards := []models.Board{{Title: "title1", Description: "desc1"}, {Title: "title2", Description: "desc2"}}
	boardRepo.EXPECT().GetUserBoards(uint(22)).Return(boards, nil)
	newBoards, err := boardUseCase.GetBoards(uint(22))
	assert.Equal(t, boards, newBoards)
	assert.Equal(t, err, nil)

	boardRepo.EXPECT().GetUserBoards(uint(22)).Return(nil, customErrors.ErrBadInputData)
	newBoards, err = boardUseCase.GetBoards(uint(22))
	assert.Equal(t, []models.Board(nil), newBoards)
	assert.Equal(t, err, customErrors.ErrBadInputData)
}

func TestGetSingleBoard(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo, commentRepo)

	board := models.Board{Title: "title2", Description: "desc2", IdB: 22, IdU: 11}
	boardRepo.EXPECT().IsAccessToBoard(uint(11), uint(22)).Return(true, nil)
	boardRepo.EXPECT().GetById(uint(22)).Return(&board, nil)
	newBoard, err := boardUseCase.GetSingleBoard(uint(22), uint(11))
	assert.Equal(t, board, newBoard)
	assert.Equal(t, err, nil)

	boardRepo.EXPECT().IsAccessToBoard(uint(11), uint(22)).Return(false, customErrors.ErrUnauthorized)
	newBoard, err = boardUseCase.GetSingleBoard(uint(22), uint(11))
	assert.Equal(t, models.Board{}, newBoard)
	assert.Equal(t, err, customErrors.ErrUnauthorized)
}

func TestCreateBoard(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo, commentRepo)

	board := models.Board{Title: "title2", Description: "desc2", IdB: 22, IdU: 11}
	moscow, _ := time.LoadLocation("Europe/Moscow")
	board.DateCreated = strconv.Itoa(time.Now().In(moscow).Day()) + " " + rtime.Now().Month().StringInCase() + " " + strconv.Itoa(time.Now().In(moscow).Year()) + ", " + time.Now().In(moscow).Format("15:04")

	boardRepo.EXPECT().Create(&board).Return(uint(22), nil)
	boardRepo.EXPECT().AppendUser(uint(22), uint(11)).Return(nil)
	boardRepo.EXPECT().GetById(uint(22)).Return(&board, nil)
	newBoard, err := boardUseCase.CreateBoard(uint(11), board)
	assert.Equal(t, &board, newBoard)
	assert.Equal(t, err, nil)

	boardRepo.EXPECT().Create(&board).Return(uint(0), customErrors.ErrNoAccess)
	newBoard2, err := boardUseCase.CreateBoard(uint(11), board)
	assert.Equal(t, (*models.Board)(nil), newBoard2)
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestRefactorBoard(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo, commentRepo)

	board := models.Board{Title: "title2", Description: "desc2", IdB: 22, IdU: 11}
	moscow, _ := time.LoadLocation("Europe/Moscow")
	board.DateCreated = strconv.Itoa(time.Now().In(moscow).Day()) + " 11" + rtime.Now().Month().StringInCase() + " " + strconv.Itoa(time.Now().In(moscow).Year()) + ", " + time.Now().In(moscow).Format("15:04")

	boardRepo.EXPECT().IsAccessToBoard(uint(11), uint(22)).Return(true, nil)
	boardRepo.EXPECT().Update(board).Return(nil)
	err := boardUseCase.RefactorBoard(uint(11), board)
	assert.Equal(t, nil, err)

	boardRepo.EXPECT().IsAccessToBoard(uint(11), uint(22)).Return(false, customErrors.ErrNoAccess)
	err = boardUseCase.RefactorBoard(uint(11), board)
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestDeleteBoard(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo, commentRepo)

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(11)).Return(true, nil)
	boardRepo.EXPECT().Delete(uint(11)).Return(nil)
	err := boardUseCase.DeleteBoard(uint(11), uint(22))
	assert.Equal(t, nil, err)

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(11)).Return(false, customErrors.ErrNoAccess)
	err = boardUseCase.DeleteBoard(uint(11), uint(22))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}

func TestGetBoard(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo, commentRepo)

	board := models.Board{Title: "title2", Description: "desc2", IdB: 22, IdU: 11}

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(11)).Return(true, nil)
	boardRepo.EXPECT().GetLists(uint(11)).Return([]models.List{}, nil)
	boardRepo.EXPECT().GetById(uint(11)).Return(&board, nil)
	boardRepo.EXPECT().GetBoardUser(uint(11)).Return([]models.User{}, nil)
	newBoard, err := boardUseCase.GetBoard(uint(11), uint(22))
	assert.Equal(t, nil, err)
	assert.Equal(t, board, newBoard)

	boardRepo.EXPECT().IsAccessToBoard(uint(22), uint(11)).Return(false, customErrors.ErrNoAccess)
	newBoard, err = boardUseCase.GetBoard(uint(11), uint(22))
	assert.Equal(t, customErrors.ErrNoAccess, err)
	assert.Equal(t, models.Board{}, newBoard)
}

func TestSaveImageBoard(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo, commentRepo)

	board := models.Board{Title: "title2", Description: "desc2", IdB: 22, IdU: 11}

	boardRepo.EXPECT().IsAccessToBoard(uint(11), uint(22)).Return(true, nil)
	boardRepo.EXPECT().SaveImage(&board, &multipart.FileHeader{}).Return(nil)
	_, err := boardUseCase.SaveImage(uint(11), &board, &multipart.FileHeader{})
	assert.Equal(t, nil, err)

	boardRepo.EXPECT().IsAccessToBoard(uint(11), uint(22)).Return(false, customErrors.ErrNoAccess)
	_, err = boardUseCase.SaveImage(uint(11), &board, &multipart.FileHeader{})
	assert.Equal(t, customErrors.ErrNoAccess, err)
}

func TestAppendUserToBoard(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo, commentRepo)

	boardRepo.EXPECT().IsAccessToBoard(uint(11), uint(22)).Return(true, nil)
	boardRepo.EXPECT().AppendUser(uint(22), uint(15))
	userRepo.EXPECT().GetUserById(uint(15)).Return(nil, customErrors.ErrUserNotFound)
	_, err := boardUseCase.AppendUserToBoard(uint(11), uint(15), uint(22))
	assert.Equal(t, customErrors.ErrUserNotFound, err)
}

func TestDeleteUserFromBoard(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	boardRepo := mock_repositories.NewMockBoardRepository(controller)
	listRepo := mock_repositories.NewMockListRepository(controller)
	taskRepo := mock_repositories.NewMockTaskRepository(controller)
	checkListRepo := mock_repositories.NewMockCheckListRepository(controller)
	userRepo := mock_repositories.NewMockUserRepository(controller)
	commentRepo := mock_repositories.NewMockCommentRepository(controller)
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo, commentRepo)

	boardRepo.EXPECT().IsAccessToBoard(uint(11), uint(22)).Return(true, nil)
	boardRepo.EXPECT().DeleteUser(uint(22), uint(15)).Return(nil)
	err := boardUseCase.DeleteUserFromBoard(uint(11), uint(15), uint(22))
	assert.Equal(t, nil, err)

	boardRepo.EXPECT().IsAccessToBoard(uint(11), uint(22)).Return(false, customErrors.ErrNoAccess)
	err = boardUseCase.DeleteUserFromBoard(uint(11), uint(15), uint(22))
	assert.Equal(t, err, customErrors.ErrNoAccess)
}
