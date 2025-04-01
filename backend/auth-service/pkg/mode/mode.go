package mode

import "os"

type mode string

const (
	debug       mode = "debug"
	development mode = "development"
	production  mode = "production"
)

type ModeManagerInterface interface {
	IsDebug() bool
	IsDevelopment() bool
	IsProduction() bool
	GetMode() string
}

type ModeManager struct {
	mode mode
}

func NewModeManager() ModeManagerInterface {
	return &ModeManager{
		mode: detectMode(),
	}
}

func detectMode() mode {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "debug" {
			return debug
		} else if arg == "development" {
			return development
		}
	}
	return production
}

// Use this only for debugging or logging purposes
func (m *ModeManager) GetMode() string {
	return string(m.mode)
}

func (m *ModeManager) IsDebug() bool {
	return m.mode == debug
}

func (m *ModeManager) IsDevelopment() bool {
	return m.mode == development
}

func (m *ModeManager) IsProduction() bool {
	return m.mode == production
}
