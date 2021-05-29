package dbMiner

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos_v2/domain/miner"
)

type ResDao struct {
	DB *gorm.DB
}

func (dao *ResDao) Create(resource miner.ResourceS) error {
	return dao.DB.Save(&resource).Error
}

func (dao *ResDao) Delete(resourceId int64) error {
	result := dao.DB.Delete(&miner.ResourceS{}, resourceId)
	return result.Error
}

func (dao *ResDao) GetById(resourceId int64) (miner.ResourceS, error) {
	resource := miner.ResourceS{}
	result := dao.DB.First(&resource, resourceId)
	return resource, result.Error
}

func (dao *ResDao) GetByTitle(resourceTitle string) ([]miner.ResourceS, error) {
	var resources []miner.ResourceS
	result := dao.DB.Where("title = ?", resourceTitle).Find(&resources)
	return resources, result.Error
}
