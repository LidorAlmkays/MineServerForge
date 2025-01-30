package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/validators"
	"github.com/go-playground/validator/v10"
)

func init() {
	validate := validator.New()
	validators.RegisterProgramModeValidator(validate)

	var mode string

	flag.StringVar(&mode, "Mode", "development", "This flags changes the program mode")
	flag.Parse()

	mode = strings.ToLower(mode)
	Flags.Mode = enums.ProgramMode(mode)
	err := validate.Struct(Flags)
	if err != nil {
		panic(err)
	}
	fmt.Println("Project Mode: " + Flags.Mode)
}

type ProgramFlags struct {
	Mode enums.ProgramMode `validate:"required,programmode"`
}

var Flags ProgramFlags
