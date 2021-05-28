package miner

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos_v2/domain/miner"
)

type ObjTypeDao struct {
	DB *gorm.DB
}

func (dao *ObjTypeDao) Create(objectType miner.ObjTypeS) error {
	return dao.DB.Save(&objectType).Error
}

func (dao *ObjTypeDao) Delete(objectTypeId int64) error {
	result := dao.DB.Delete(&miner.ObjTypeS{}, objectTypeId)
	return result.Error
}

func (dao *ObjTypeDao) GetById(objectTypeId int64) (miner.ObjTypeS, error) {
	objectType := miner.ObjTypeS{}
	result := dao.DB.First(&objectType, objectTypeId)
	return objectType, result.Error
}

func (dao *ObjTypeDao) GetByTitle(objectTypeTitle string) ([]miner.ObjTypeS, error) {
	var objectTypes []miner.ObjTypeS
	result := dao.DB.Where("title = ?", objectTypeTitle).Find(&objectTypes)
	return objectTypes, result.Error
}
