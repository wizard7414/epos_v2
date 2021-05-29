package database

import (
	"github.com/wizard7414/epos_v2/database/dbMiner"
	"github.com/wizard7414/epos_v2/database/general"
)

type EposDatabase struct {
	DbMiner   dbMiner.DbMiner
	DbGeneral general.DbGeneral
}

func (db *EposDatabase) CloseMiner() {
	err := db.DbMiner.CloseDb()
	if err != nil {
		panic(err)
	}
}

func (db *EposDatabase) CloseGeneral() {
	//err := db.DbGeneral.InitDb(dbPath)
	/*if err != nil {
		panic(err)
	}*/
}
