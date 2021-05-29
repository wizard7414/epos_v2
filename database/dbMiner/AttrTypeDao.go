package dbMiner

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos_v2/domain/miner"
)

type AttrTypeDao struct {
	DB *gorm.DB
}

func (dao *AttrTypeDao) Create(attributeType miner.AttrTypeS) error {
	return dao.DB.Save(&attributeType).Error
}

func (dao *AttrTypeDao) Delete(attributeTypeId int64) error {
	result := dao.DB.Delete(&miner.AttrTypeS{}, attributeTypeId)
	return result.Error
}

func (dao *AttrTypeDao) GetById(attributeTypeId int64) (miner.AttrTypeS, error) {
	attributeType := miner.AttrTypeS{}
	result := dao.DB.First(&attributeType, attributeTypeId)
	return attributeType, result.Error
}

func (dao *AttrTypeDao) GetByTitle(attributeTypeTitle string) ([]miner.AttrTypeS, error) {
	var attributeTypes []miner.AttrTypeS
	result := dao.DB.Where("title = ?", attributeTypeTitle).Find(&attributeTypes)
	return attributeTypes, result.Error
}
