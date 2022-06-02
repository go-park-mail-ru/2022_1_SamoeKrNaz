package impl

import (
	"PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/models"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"gorm.io/gorm"
	"image"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

const filePathBoards = "img_boards/"

type BoardRepositoryImpl struct {
	db *gorm.DB
}

func MakeBoardRepository(db *gorm.DB) repositories.BoardRepository {
	return &BoardRepositoryImpl{db: db}
}

func (boardRepository *BoardRepositoryImpl) Create(board *models.Board) (uint, error) {
	err := boardRepository.db.Create(board).Error
	return board.IdB, err
}

func (boardRepository *BoardRepositoryImpl) AppendUser(boardId uint, IdU uint) error {
	user := new(models.User)
	result := boardRepository.db.Find(user, IdU)
	if result.RowsAffected == 0 {
		return customErrors.ErrUserNotFound
	} else if result.Error != nil {
		return result.Error
	}
	err := boardRepository.db.Model(&models.Board{IdB: boardId}).Association("Users").Append(user)
	return err
}

func (boardRepository *BoardRepositoryImpl) GetLists(IdB uint) ([]models.List, error) {
	lists := new([]models.List)
	result := boardRepository.db.Where("id_b = ?", IdB).Order("position").Find(lists)
	return *lists, result.Error
}

func (boardRepository *BoardRepositoryImpl) Update(board models.Board) error {
	// возьмем из бд текущую запись по айдишнику
	currentData, err := boardRepository.GetById(board.IdB)
	// обработка ошибки при взятии
	if err != nil {
		return err
	}
	// ищем, какое поле поменялось
	if currentData.Title != board.Title && board.Title != "" {
		currentData.Title = board.Title
	}
	if currentData.Description != board.Description && board.Description != "" {
		currentData.Description = board.Description
	}
	//сохраняем новую структуру
	return boardRepository.db.Save(currentData).Error
}

func (boardRepository *BoardRepositoryImpl) Delete(IdB uint) error {
	err := boardRepository.db.Model(&models.Board{IdB: IdB}).Association("Users").Clear()
	if err != nil {
		return err
	}
	return boardRepository.db.Delete(&models.Board{}, IdB).Error
}

func (boardRepository *BoardRepositoryImpl) GetUserBoards(IdU uint) ([]models.Board, error) {
	boards := new([]models.Board)
	err := boardRepository.db.Model(&models.User{IdU: IdU}).Association("Boards").Find(boards)
	if err != nil {
		return nil, err
	}
	return *boards, nil
}

func (boardRepository *BoardRepositoryImpl) GetById(IdB uint) (*models.Board, error) {
	// указатель на структуру, которую вернем
	board := new(models.Board)
	result := boardRepository.db.Find(board, IdB)
	// если выборка в 0 строк, то такой доски нет
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrBoardNotFound
	} else if result.Error != nil {
		// если произошла ошибка при выборке
		return nil, result.Error
	}
	// иначе вернем доску
	return board, nil
}

func (boardRepository *BoardRepositoryImpl) IsAccessToBoard(IdU uint, IdB uint) (bool, error) {
	user := new(models.User)
	err := boardRepository.db.Model(&models.Board{IdB: IdB}).Where("id_u = ?", IdU).Association("Users").Find(user)
	if err != nil {
		return false, err
	} else if user.IdU == 0 {
		return false, nil
	}
	return true, nil
}

func (boardRepository *BoardRepositoryImpl) SaveImage(board *models.Board, header *multipart.FileHeader) error {
	if board.ImgDesk != "" {
		currentData, err := boardRepository.GetById(board.IdB)
		if err != nil {
			return err
		}

		fileName := strings.Join([]string{filePathBoards, strconv.Itoa(int(currentData.IdB)), ".webp"}, "")
		output, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer output.Close()

		openFile, err := header.Open()
		if err != nil {
			return err
		}

		img, _, err := image.Decode(openFile)
		if err != nil {
			return err
		}

		options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
		if err != nil {
			return err
		}

		err = webp.Encode(output, img, options)
		if err != nil {
			return err
		}

		currentData.ImgDesk = fileName
		return boardRepository.db.Save(currentData).Error
	}
	return nil
}

func (boardRepository *BoardRepositoryImpl) GetBoardUser(IdB uint) ([]models.User, error) {
	users := new([]models.User)
	err := boardRepository.db.Model(&models.Board{IdB: IdB}).Association("Users").Find(users)
	if err != nil {
		return nil, err
	}
	return *users, nil
}

func (boardRepository *BoardRepositoryImpl) DeleteUser(boardId uint, userId uint) error {
	user := new(models.User)
	result := boardRepository.db.Find(user, userId)
	if result.RowsAffected == 0 {
		return customErrors.ErrUserNotFound
	} else if result.Error != nil {
		return result.Error
	}
	err := boardRepository.db.Model(&models.Board{IdB: boardId}).Association("Users").Delete(user)
	return err
}

func (boardRepository *BoardRepositoryImpl) GetByLink(link string) (*models.Board, error) {
	// указатель на структуру, которую вернем
	board := new(models.Board)
	result := boardRepository.db.Where("link = ?", link).Find(board)
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrBoardNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}
	return board, nil
}
