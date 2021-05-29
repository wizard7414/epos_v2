package main

import (
	"fmt"
	"github.com/wizard7414/epos_v2/database/dbMiner"
	"github.com/wizard7414/epos_v2/parser/deviant"
	"github.com/wizard7414/epos_v2/utils"
)

func main() {
	configuration := utils.GetConfig("config/deviant-conf.json")

	/*path := "/home/wizard/yDisk/#Epos/Artist/7/7zaki"

	files, err := utility.GetFiles(path)

	if err != nil {
		panic(err)
	}

	println(len(files))*/

	base := dbMiner.DbMiner{
		AttributeCode:   dbMiner.AttrCodeDao{},
		AttributeType:   dbMiner.AttrTypeDao{},
		Attribute:       dbMiner.AttrDao{},
		Resource:        dbMiner.ResDao{},
		Source:          dbMiner.SrcDao{},
		Entry:           dbMiner.EntDao{},
		ObjectType:      dbMiner.ObjTypeDao{},
		Object:          dbMiner.ObjDao{},
		ObjectAttribute: dbMiner.ObjAttrDao{},
	}

	err := base.InitDb(configuration.DbMinerPath)
	if err != nil {
		panic("Unable to start test set!")
	}
	defer base.CloseDb()

	completesObjects, rejectedObjects := deviant.ParseStage(&base, configuration.PathToBuffer, configuration.AuthHeader)
	if len(completesObjects) > 0 && len(rejectedObjects) == 0 {

		if configuration.GraphicsPath != "" {
			if configuration.GraphVpnSave {
				deviant.SaveGraphStageVpn(completesObjects, configuration.GraphicsPath)
			} else {
				deviant.SaveGraphStage(completesObjects, configuration.GraphicsPath)
			}
		}

		fmt.Println("==============================================================")

		if configuration.DbMinerPath != "" {
			deviant.SaveObjectsStage(&base, completesObjects)
		}

	}
}
