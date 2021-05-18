package utils

import "testing"

func Test_GetConfig(t *testing.T) {
	conf := GetConfig("../config/test-conf.json")

	if conf.ParseLimit != 50 || conf.GraphVpnSave != false || conf.PathToBuffer != "/testPath" ||
		conf.DbMinerPath != "/testDbPath" || conf.GraphicsPath != "/testPath" || conf.AuthHeader != "test-header" ||
		conf.DbGeneralPath != "/testDbPath" {
		t.Fail()
	}
}
