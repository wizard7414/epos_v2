package dbMiner

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos_v2/domain/miner"
)

type ObjAttrDao struct {
	DB *gorm.DB
}

func (dao *ObjAttrDao) Create(objectAttribute miner.ObjAttrS) error {
	return dao.DB.Save(&objectAttribute).Error
}

func (dao *ObjAttrDao) Delete(objectId int64, attributeId int64) error {
	objectAttribute := miner.ObjAttrS{}
	result := dao.DB.Where("object = ? AND attribute = ?", objectId, attributeId).Delete(&objectAttribute)
	return result.Error
}

func (dao *ObjAttrDao) GetById(objectId int64, attributeId int64) (miner.ObjAttrS, error) {
	objectAttribute := miner.ObjAttrS{}
	result := dao.DB.Where("object = ? AND attribute = ?", objectId, attributeId).First(&objectAttribute)
	return objectAttribute, result.Error
}

func (dao *ObjAttrDao) GetByObject(objectId int64) ([]miner.ObjAttrS, error) {
	var objectAttributes []miner.ObjAttrS
	result := dao.DB.Where("object = ?", objectId).Find(&objectAttributes)
	return objectAttributes, result.Error
}

func (dao *ObjAttrDao) GetByAttribute(attributeId int64) ([]miner.ObjAttrS, error) {
	var objectAttributes []miner.ObjAttrS
	result := dao.DB.Where("attribute = ?", attributeId).Find(&objectAttributes)
	return objectAttributes, result.Error
}
