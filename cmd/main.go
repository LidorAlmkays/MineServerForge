package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/LidorAlmkays/MineServerForge/utils/validators"
	"github.com/go-playground/validator"
)



type ProgramFlags struct {
	Mode string `validate:"required,programmode"`
}

var programFlags ProgramFlags

func init() {
	flag.StringVar(&programFlags.Mode, "Mode", "development", "This flags changes the program mode")
	flag.Parse()
	programFlags.Mode = strings.ToLower(programFlags.Mode)
	validate := validator.New()
	validate.RegisterValidation("programmode", validators.ProgramModeValidator)
	err := validate.Struct(programFlags)
	if err != nil {
		panic(err)
	}
}

func main() {
	err := setUp()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	os.Exit(0)
}


func setUp() error {
	// ctx := context.Background()
	// var err error

	return nil
}


