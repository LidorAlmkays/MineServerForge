package enums

type ProgramMode string

const (
	DevelopmentMode ProgramMode = "development"
	ProductionMode  ProgramMode = "production"
)

// IsValid checks if the ProgramMode is valid
func (w ProgramMode) IsValid() bool {
	switch w {
	case DevelopmentMode, ProductionMode:
		return true
	default:
		return false
	}
}
