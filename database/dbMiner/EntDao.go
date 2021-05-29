package dbMiner

import (
	"github.com/jinzhu/gorm"
	"github.com/wizard7414/epos_v2/domain/miner"
)

type EntDao struct {
	DB *gorm.DB
}

func (dao *EntDao) Create(entry miner.EntryS) error {
	return dao.DB.Save(&entry).Error
}

func (dao *EntDao) Delete(entryId int64) error {
	result := dao.DB.Delete(&miner.EntryS{}, entryId)
	return result.Error
}

func (dao *EntDao) GetById(entryId int64) (miner.EntryS, error) {
	entry := miner.EntryS{}
	result := dao.DB.First(&entry, entryId)
	return entry, result.Error
}

func (dao *EntDao) GetByTitle(entryTitle string) ([]miner.EntryS, error) {
	var entries []miner.EntryS
	result := dao.DB.Where("title = ?", entryTitle).Find(&entries)
	return entries, result.Error
}
