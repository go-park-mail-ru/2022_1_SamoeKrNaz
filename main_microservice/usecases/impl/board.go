package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/main_microservice/usecases"
	"PLANEXA_backend/models"
	"github.com/google/uuid"
	rtime "github.com/ivahaev/russian-time"
	"github.com/microcosm-cc/bluemonday"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type BoardUseCaseImpl struct {
	repBoard     repositories.BoardRepository
	repList      repositories.ListRepository
	repTask      repositories.TaskRepository
	repCheckList repositories.CheckListRepository
	repUser      repositories.UserRepository
	repComment   repositories.CommentRepository
}

func MakeBoardUsecase(repBoard_ repositories.BoardRepository, repList_ repositories.ListRepository,
	repTask_ repositories.TaskRepository, repCheckList_ repositories.CheckListRepository,
	repUser_ repositories.UserRepository, repComment_ repositories.CommentRepository) usecases.BoardUseCase {
	return &BoardUseCaseImpl{repBoard: repBoard_, repList: repList_,
		repTask: repTask_, repCheckList: repCheckList_, repUser: repUser_,
		repComment: repComment_}
}

func (boardUseCase *BoardUseCaseImpl) GetBoards(userId uint) ([]models.Board, error) {
	// достаю из БД доски по userId
	boards, err := boardUseCase.repBoard.GetUserBoards(userId)
	if err != nil {
		return nil, err
	}
	sanitizer := bluemonday.UGCPolicy()
	for _, board := range boards {
		board.DateCreated = sanitizer.Sanitize(board.DateCreated)
		board.Title = sanitizer.Sanitize(board.Title)
		board.Description = sanitizer.Sanitize(board.Description)
		board.ImgDesk = sanitizer.Sanitize(board.ImgDesk)
	}
	return boards, nil
}

func (boardUseCase *BoardUseCaseImpl) GetSingleBoard(boardId uint, userId uint) (models.Board, error) {
	//проверить может ли юзер смотреть эту доску
	// вызываю из бд получение доски
	isAccess, err := boardUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return models.Board{}, err
	} else if !isAccess {
		return models.Board{}, customErrors.ErrNoAccess
	}

	board, err := boardUseCase.repBoard.GetById(boardId)
	if err != nil {
		return models.Board{}, err
	}
	sanitizer := bluemonday.UGCPolicy()
	board.DateCreated = sanitizer.Sanitize(board.DateCreated)
	board.Title = sanitizer.Sanitize(board.Title)
	board.Description = sanitizer.Sanitize(board.Description)
	board.ImgDesk = sanitizer.Sanitize(board.ImgDesk)
	return *board, nil
}

func (boardUseCase *BoardUseCaseImpl) CreateBoard(userId uint, board models.Board) (*models.Board, error) {
	// добавляю в бд такую доску с привязкой к данному юзеру
	moscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, err
	}
	board.DateCreated = strconv.Itoa(time.Now().In(moscow).Day()) + " " + rtime.Now().Month().StringInCase() + " " + strconv.Itoa(time.Now().In(moscow).Year()) + ", " + time.Now().In(moscow).Format("15:04")
	board.IdU = userId
	board.Link = uuid.NewString()
	boardId, err := boardUseCase.repBoard.Create(&board)
	if err != nil {
		return nil, err
	}
	err = boardUseCase.repBoard.AppendUser(board.IdB, userId)
	if err != nil {
		return nil, err
	}
	createdBoard, err := boardUseCase.repBoard.GetById(boardId)
	return createdBoard, err
}

func (boardUseCase *BoardUseCaseImpl) RefactorBoard(userId uint, board models.Board) error {
	// проверяю есть ли доска с таким айди и может ли юзер её редачить
	//вызываю репозиторий для обновления доски
	isAccess, err := boardUseCase.repBoard.IsAccessToBoard(userId, board.IdB)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	err = boardUseCase.repBoard.Update(board)
	return err
}

func (boardUseCase *BoardUseCaseImpl) DeleteBoard(boardId uint, userId uint) error {
	// проверяю есть ли такая доска и может ли юзер редачить её
	// удаляю из бд
	isAccess, err := boardUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	err = boardUseCase.repBoard.Delete(boardId)
	return err
}

func (boardUseCase *BoardUseCaseImpl) GetBoard(boardId, userId uint) (models.Board, error) {
	isAccess, err := boardUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return models.Board{}, err
	} else if !isAccess {
		return models.Board{}, customErrors.ErrNoAccess
	}
	lists, err := boardUseCase.repBoard.GetLists(boardId)
	if err != nil {
		return models.Board{}, err
	}
	for i, list := range lists {
		tasks, err := boardUseCase.repList.GetTasks(list.IdL)
		if err != nil {
			return models.Board{}, err
		}
		for j, task := range *tasks {
			appendedUsers, err := boardUseCase.repTask.GetTaskUser(task.IdT)
			if err != nil {
				return models.Board{}, err
			}
			(*tasks)[j].Users = *appendedUsers
			checkLists, err := boardUseCase.repTask.GetCheckLists(task.IdT)
			if err != nil {
				return models.Board{}, err
			}
			for k, checkList := range *checkLists {
				checkListItem, err := boardUseCase.repCheckList.GetCheckListItems(checkList.IdCl)
				if err != nil {
					return models.Board{}, err
				}
				(*checkLists)[k].CheckListItems = *checkListItem
			}
			comments, err := boardUseCase.repComment.GetComments(task.IdT)
			if err != nil {
				return models.Board{}, err
			}
			for i, comment := range *comments {
				userComment, err := boardUseCase.repUser.GetUserById(comment.IdU)
				if err != nil {
					return models.Board{}, err
				}
				(*comments)[i].User = *userComment
			}
			attachments, err := boardUseCase.repTask.GetAttachments(task.IdT)
			if err != nil {
				return models.Board{}, err
			}
			(*tasks)[j].Comments = *comments
			(*tasks)[j].CheckLists = *checkLists
			(*tasks)[j].Attachments = *attachments
		}
		lists[i].Tasks = *tasks
	}
	board, err := boardUseCase.repBoard.GetById(boardId)
	if err != nil {
		return models.Board{}, err
	}
	appendedUsers, err := boardUseCase.repBoard.GetBoardUser(boardId)
	if err != nil {
		return models.Board{}, err
	}
	for _, user := range appendedUsers {
		user.Password = ""
	}
	board.Users = appendedUsers
	if err != nil {
		return models.Board{}, err
	}
	sanitizer := bluemonday.UGCPolicy()
	board.DateCreated = sanitizer.Sanitize(board.DateCreated)
	board.Title = sanitizer.Sanitize(board.Title)
	board.Description = sanitizer.Sanitize(board.Description)
	board.ImgDesk = sanitizer.Sanitize(board.ImgDesk)
	board.Lists = lists
	return *board, nil
}

func (boardUseCase *BoardUseCaseImpl) SaveImage(userId uint, board *models.Board, header *multipart.FileHeader) (string, error) {
	isAccess, err := boardUseCase.repBoard.IsAccessToBoard(userId, board.IdB)
	if err != nil {
		return "", err
	} else if !isAccess {
		return "", customErrors.ErrNoAccess
	}
	err = boardUseCase.repBoard.SaveImage(board, header)
	return strings.Join([]string{strconv.Itoa(int(board.IdB)), ".webp"}, ""), err
}

func (boardUseCase *BoardUseCaseImpl) AppendUserToBoard(userId uint, appendedUserId uint, boardId uint) (models.User, error) {
	isAccess, err := boardUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return models.User{}, err
	} else if !isAccess {
		return models.User{}, customErrors.ErrNoAccess
	}
	err = boardUseCase.repBoard.AppendUser(boardId, appendedUserId)
	if err != nil {
		return models.User{}, err
	}
	user, err := boardUseCase.repUser.GetUserById(appendedUserId)
	if err != nil {
		return models.User{}, err
	}
	return *user, err
}

func (boardUseCase *BoardUseCaseImpl) DeleteUserFromBoard(userId uint, deletedUserId uint, boardId uint) error {
	isAccess, err := boardUseCase.repBoard.IsAccessToBoard(userId, boardId)
	if err != nil {
		return err
	} else if !isAccess {
		return customErrors.ErrNoAccess
	}
	err = boardUseCase.repBoard.DeleteUser(boardId, deletedUserId)
	if err != nil {
		return err
	}
	return err
}

func (boardUseCase *BoardUseCaseImpl) AppendUserByLink(userId uint, link string) (*models.Board, error) {
	board, err := boardUseCase.repBoard.GetByLink(link)
	if err != nil {
		return nil, err
	}
	isAccess, err := boardUseCase.repBoard.IsAccessToBoard(userId, board.IdB)
	if isAccess {
		return nil, customErrors.ErrAlreadyAppended
	} else if err != nil {
		return nil, err
	}
	err = boardUseCase.repBoard.AppendUser(board.IdB, userId)
	if err != nil {
		return nil, err
	}
	return board, err
}
