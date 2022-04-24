package impl

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"PLANEXA_backend/repositories"
	"gorm.io/gorm"
)

type CheckListItemRepositoryImpl struct {
	db *gorm.DB
}

func MakeCheckListItemRepository(db *gorm.DB) repositories.CheckListItemRepository {
	return &CheckListItemRepositoryImpl{db: db}
}

func (checkListItemRepository *CheckListItemRepositoryImpl) Create(checkListItem *models.CheckListItem) (uint, error) {
	err := checkListItemRepository.db.Create(checkListItem).Error
	return checkListItem.IdClIt, err
}

func (checkListItemRepository *CheckListItemRepositoryImpl) GetById(IdCl uint) (*models.CheckListItem, error) {
	checkListItem := new(models.CheckListItem)
	result := checkListItemRepository.db.Find(checkListItem, IdCl)
	if result.RowsAffected == 0 {
		return nil, customErrors.ErrBoardNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}
	return checkListItem, nil
}

func (checkListItemRepository *CheckListItemRepositoryImpl) Update(checkListItem models.CheckListItem) error {
	currentData, err := checkListItemRepository.GetById(checkListItem.IdClIt)
	if err != nil {
		return err
	}
	if currentData.Description != checkListItem.Description && checkListItem.Description != "" {
		currentData.Description = checkListItem.Description
	}
	if currentData.IsReady != checkListItem.IsReady && checkListItem.IsReady != false {
		currentData.IsReady = checkListItem.IsReady
	}
	return checkListItemRepository.db.Save(currentData).Error
}

func (checkListItemRepository *CheckListItemRepositoryImpl) Delete(IdClIt uint) error {
	return checkListItemRepository.db.Delete(&models.CheckListItem{}, IdClIt).Error
}

func (checkListItemRepository *CheckListItemRepositoryImpl) GetCheckListItems(IdCl uint) ([]models.CheckListItem, error) {
	checkListItems := new([]models.CheckListItem)
	result := checkListItemRepository.db.Where("id_cl = ?", IdCl).Find(checkListItems)
	return *checkListItems, result.Error
}
