package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories/mocks"
	"github.com/golang/mock/gomock"
	rtime "github.com/ivahaev/russian-time"
	"github.com/stretchr/testify/assert"
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
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo)

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
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo)

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
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo)

	board := models.Board{Title: "title2", Description: "desc2", IdB: 22, IdU: 11}
	moscow, _ := time.LoadLocation("Europe/Moscow")
	board.DateCreated = strconv.Itoa(time.Now().In(moscow).Day()) + " " + rtime.Now().Month().StringInCase() + " " + strconv.Itoa(time.Now().In(moscow).Year()) + ", " + strconv.Itoa(time.Now().In(moscow).Hour()) + ":" + strconv.Itoa(time.Now().In(moscow).Minute())

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
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo)

	board := models.Board{Title: "title2", Description: "desc2", IdB: 22, IdU: 11}
	moscow, _ := time.LoadLocation("Europe/Moscow")
	board.DateCreated = strconv.Itoa(time.Now().In(moscow).Day()) + " " + rtime.Now().Month().StringInCase() + " " + strconv.Itoa(time.Now().In(moscow).Year()) + ", " + strconv.Itoa(time.Now().In(moscow).Hour()) + ":" + strconv.Itoa(time.Now().In(moscow).Minute())

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
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo)

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
	boardUseCase := MakeBoardUsecase(boardRepo, listRepo, taskRepo, checkListRepo, userRepo)

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
