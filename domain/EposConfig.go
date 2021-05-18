package domain

type EposConfig struct {
	PathToBuffer string
	ParseLimit   int
	AuthHeader   string
	SaveToDb     bool
	GraphicsPath string
	GraphVpnSave bool
}
