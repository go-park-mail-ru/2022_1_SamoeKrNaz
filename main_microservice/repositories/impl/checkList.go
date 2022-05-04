package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/main_microservice/repositories"
	"PLANEXA_backend/models"
	"gorm.io/gorm"
)

type CheckListRepositoryImpl struct {
	db *gorm.DB
}

func MakeCheckListRepository(db *gorm.DB) repositories.CheckListRepository {
	return &CheckListRepositoryImpl{db: db}
}

func (checkListRepository *CheckListRepositoryImpl) Create(checkList *models.CheckList) (uint, error) {
	err := checkListRepository.db.Create(checkList).Error
	return checkList.IdCl, err
}

func (checkListRepository *CheckListRepositoryImpl) GetById(IdCl uint) (*models.CheckList, error) {
	checkList := new(models.CheckList)
	result := checkListRepository.db.Find(checkList, IdCl)
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrCheckListNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}
	return checkList, nil
}

func (checkListRepository *CheckListRepositoryImpl) Update(checkList models.CheckList) error {
	currentData, err := checkListRepository.GetById(checkList.IdCl)
	if err != nil {
		return err
	}
	if currentData.Title != checkList.Title && checkList.Title != "" {
		currentData.Title = checkList.Title
	}
	return checkListRepository.db.Save(currentData).Error
}

func (checkListRepository *CheckListRepositoryImpl) Delete(IdCl uint) error {
	return checkListRepository.db.Delete(&models.CheckList{}, IdCl).Error
}

func (checkListRepository *CheckListRepositoryImpl) GetCheckListItems(IdCl uint) (*[]models.CheckListItem, error) {
	checkListItems := new([]models.CheckListItem)
	err := checkListRepository.db.Where("id_cl = ?", IdCl).Order("id_cl_it").Find(checkListItems).Error
	return checkListItems, err
}
