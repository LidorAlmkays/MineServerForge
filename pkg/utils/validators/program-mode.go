package validators

import (
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
	"github.com/go-playground/validator/v10"
)

func RegisterProgramModeValidator(validate *validator.Validate) {
	// Custom validation for optional struct pointers
	validate.RegisterValidation("programmode", func(fl validator.FieldLevel) bool {
		var value enums.ProgramMode = enums.ProgramMode(fl.Field().String())
		return value.IsValid()
	})
}
