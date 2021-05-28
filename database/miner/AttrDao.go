package miner

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos_v2/domain/miner"
)

type AttrDao struct {
	DB *gorm.DB
}

func (dao *AttrDao) Create(attribute miner.AttrS) error {
	return dao.DB.Save(&attribute).Error
}

func (dao *AttrDao) Delete(attributeId int64) error {
	result := dao.DB.Delete(&miner.AttrS{}, attributeId)
	return result.Error
}

func (dao *AttrDao) GetById(attributeId int64) (miner.AttrS, error) {
	attribute := miner.AttrS{}
	result := dao.DB.First(&attribute, attributeId)
	return attribute, result.Error
}

func (dao *AttrDao) GetByCode(attributeCode int64) ([]miner.AttrS, error) {
	var attributes []miner.AttrS
	result := dao.DB.Where("code = ?", attributeCode).Find(&attributes)
	return attributes, result.Error
}
