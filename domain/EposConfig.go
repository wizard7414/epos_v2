package domain

type EposConfig struct {
	PathToBuffer  string
	ParseLimit    int
	AuthHeader    string
	DbMinerPath   string
	DbGeneralPath string
	GraphicsPath  string
	GraphVpnSave  bool
}
